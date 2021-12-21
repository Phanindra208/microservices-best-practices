package tracing

import (
	"net/http"

	"github.com/opentracing/opentracing-go"
)

// HTTPClient is a wrapper for executing the request in a span.
type HTTPClient struct {
	http.Client
}

// Do wraps (http.Client).Do func.
func (client *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	span, ctx := opentracing.StartSpanFromContext(req.Context(), "")
	defer span.Finish()

	configureSpan(ctx, span, DirectionOutgoing)

	req = InjectSpan(req.WithContext(ctx))
	dumpHTTPRequest(req)
	resp, err := client.Client.Do(req)
	if err != nil {
		span.SetTag("error", err.Error())
	}
	if resp != nil {
		dumpHTTPResponse(resp)
		span.SetTag("status_code", resp.StatusCode)
	}

	return resp, err
}
