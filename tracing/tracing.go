package tracing

import (
	"fmt"
	"html"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/iAmPlus/microservice/config"
	"github.com/iAmPlus/microservice/log"
	"github.com/justinas/alice"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	zipkin "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go-opentracing/thrift/gen-go/zipkincore"
	"go.uber.org/zap"
)

// InitTracer creates a new opentracing tracer and overrides the global tracer.
func InitTracer(zipkinHTTPEndpoint string, hostPort string, debug bool, serviceName string) {
	log := log.Sugar()

	collector, err := zipkin.NewHTTPCollector(zipkinHTTPEndpoint, newErrorLogger())
	if err != nil {
		log.Errorf("unable to create Zipkin HTTP collector: %v", err)
		return
	}
	recorder := zipkin.NewRecorder(collector, debug, hostPort, serviceName)

	tracer, err := zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(true),
		zipkin.TraceID128Bit(true),
	)
	if err != nil {
		log.Errorf("unable to create Zipkin tracer: %v", err)
		return
	}

	// Override the global tracer.
	opentracing.SetGlobalTracer(tracer)
}

type zipkinErrorLogger struct {
	logger *zap.SugaredLogger
}

func (zel *zipkinErrorLogger) Log(args ...interface{}) error {
	if config.Vars.LogZipkinErrors {
		zel.logger.Debug(args...)
	}
	return nil
}

func newErrorLogger() zipkin.HTTPOption {
	return zipkin.HTTPLogger(&zipkinErrorLogger{logger: log.Sugar()})
}

type statusResponseWriter struct {
	status   int
	response []byte
	http.ResponseWriter
}

func (srw *statusResponseWriter) WriteHeader(statusCode int) {
	srw.status = statusCode
	srw.ResponseWriter.WriteHeader(statusCode)
}

func (srw *statusResponseWriter) Write(b []byte) (int, error) {
	srw.response = b
	return srw.ResponseWriter.Write(b)
}

// GetMiddleware is a tracing middleware.
func GetMiddleware() alice.Constructor {
	tracer := opentracing.GlobalTracer()

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path := html.EscapeString(r.URL.Path)
			if !(strings.Contains(path, "liveness") || strings.Contains(path, "readiness")) {
				dumpHTTPRequest(r)
				var span opentracing.Span
				span, r = incomingHTTPRequest(tracer, r)
				srw := &statusResponseWriter{status: 200, ResponseWriter: w}
				h.ServeHTTP(srw, r)
				if span != nil {
					span.SetTag("status_code", srw.status)
					span.Finish()
				}
				if len(srw.response) > 0 {
					logResponseBytes(srw.response)
				}
			} else {
				h.ServeHTTP(w, r)
			}
		})
	}
}

func dumpHTTPRequest(r *http.Request) {
	if !config.Vars.EnableZipkinLogs {
		return
	}

	l := log.Sugar()
	dump, _ := httputil.DumpRequest(r, true)
	l.Debug("HTTP request: ", makeSingleLineHTTP(dump))
}

func dumpHTTPResponse(r *http.Response) {
	if !config.Vars.EnableZipkinLogs {
		return
	}

	l := log.Sugar()
	dump, _ := httputil.DumpResponse(r, true)
	l.Debug("HTTP response: ", makeSingleLineHTTP(dump))
}

func logResponseBytes(b []byte) {
	if !config.Vars.EnableZipkinLogs {
		return
	}

	l := log.Sugar()
	l.Debug("HTTP response: ", makeSingleLineHTTP(b))
}

func makeSingleLineHTTP(b []byte) string {
	if config.Vars.TraceMaxBytes > 0 && len(b) > config.Vars.TraceMaxBytes {
		b = b[:config.Vars.TraceMaxBytes]
	}
	return strings.Join(strings.Fields(string(b)), " ")
}

// InjectSpan returns a new request which contains the OpenTracing Span in the
// context. If no such Span can be found, then it is a noop.
func InjectSpan(req *http.Request) *http.Request {
	span := opentracing.SpanFromContext(req.Context())
	if span == nil {
		return req
	}

	// We are going to use this span in a client request, so mark as such.
	ext.SpanKindRPCClient.Set(span)

	// Add some standard OpenTracing tags, useful in an HTTP request.
	ext.HTTPMethod.Set(span, req.Method)
	span.SetTag(zipkincore.HTTP_HOST, req.URL.Host)
	span.SetTag(zipkincore.HTTP_PATH, req.URL.Path)
	ext.HTTPUrl.Set(
		span,
		fmt.Sprintf("%s://%s%s", req.URL.Scheme, req.URL.Host, req.URL.Path),
	)

	// Add information on the peer service we're about to contact.
	if host, portString, err := net.SplitHostPort(req.URL.Host); err == nil {
		ext.PeerHostname.Set(span, host)
		if port, err := strconv.Atoi(portString); err != nil {
			ext.PeerPort.Set(span, uint16(port))
		}
	} else {
		ext.PeerHostname.Set(span, req.URL.Host)
		ext.PeerPort.Set(span, 80)
	}

	// Inject the Span context into the outgoing HTTP Request.
	if err := opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.TextMap,
		opentracing.HTTPHeadersCarrier(req.Header),
	); err != nil {
		log.Sugar().Errorf("failed to inject span: %v", err)
	}

	return req
}

func incomingHTTPRequest(tracer opentracing.Tracer, req *http.Request) (opentracing.Span, *http.Request) {
	wireContext, err := tracer.Extract(
		opentracing.TextMap,
		opentracing.HTTPHeadersCarrier(req.Header),
	)
	if err != nil {
		log.Sugar().Errorf("failed to extract span: %v", err)
	}

	span := tracer.StartSpan(req.URL.Path, ext.RPCServerOption(wireContext))

	ctx := opentracing.ContextWithSpan(req.Context(), span)
	route := middleware.MatchedRouteFrom(req)
	ctx = Context(ctx).With(API, config.Vars.ZipkinServiceName).With(Name, route.Operation.ID).Ctx()
	configureSpan(ctx, span, DirectionIncoming)

	ext.HTTPMethod.Set(span, req.Method)
	span.SetTag(zipkincore.HTTP_PATH, req.URL.Path)

	return span, req.WithContext(ctx)
}
