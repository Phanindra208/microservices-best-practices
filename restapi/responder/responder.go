package responder

import (
	"net/http"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/iAmPlus/microservice/log"
)

// Responder implements middleware.Responder and allows us to
// build responses easily.
type Responder struct {
	req    *http.Request
	status int
	body   interface{}
}

// New creates a new responder.
func New(req *http.Request) (resp *Responder) {
	resp = &Responder{req: req}
	return resp
}

// WriteResponse writes HTTP response.
func (resp *Responder) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	if resp.status == 0 {
		resp.status = 200
	}

	// Put other logic before writing.

	resp.write(rw, producer)
}

// write writes response:
//   1) Set header fields.
//   2) Write header with status.
//   3) Write body.
func (resp *Responder) write(rw http.ResponseWriter, producer runtime.Producer) {
	// If response body is nil, remove Content-Type header field only.
	if resp.body == nil {
		rw.Header().Del(runtime.HeaderContentType)
	}

	rw.WriteHeader(resp.status)

	if resp.body != nil {
		resp.fixNull()

		var err error
		if b, ok := resp.body.([]byte); ok {
			_, err = rw.Write(b)
		} else {
			err = producer.Produce(rw, resp.body)
		}

		if err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// fixNull makes sure that we don't marshal to "null".
func (resp *Responder) fixNull() {
	v := reflect.ValueOf(resp.body)
	if v.Kind() == reflect.Slice && v.Len() == 0 {
		resp.body = []int{} // Any array which is marshalled to []
	}
}

// Status sets HTTP status code.
func (resp *Responder) Status(status int) *Responder {
	resp.status = status
	return resp
}

// Body sets response body to respond with.
func (resp *Responder) Body(body interface{}) *Responder {
	resp.body = body
	return resp
}

// LogErr logs internal error.
// TODO: Use customizable messages?
func (resp *Responder) LogErr(err error) *Responder {
	l := log.Sugar()

	switch {
	case resp.status >= 500:
		l.Errorw(
			err.Error(),
			"correlation-id", resp.req.Header.Get("X-Correlation-ID"),
		)

	case resp.status >= 400:
		l.Warnw(
			err.Error(),
			"correlation-id", resp.req.Header.Get("X-Correlation-ID"),
		)

	default:
		l.Infow(
			err.Error(),
			"correlation-id", resp.req.Header.Get("X-Correlation-ID"),
		)
	}

	return resp
}

type customErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

// Error prepares response body.
func (resp *Responder) Error(code int, msg ...string) *Responder {
	body := &customErrorResponse{Code: code}
	if msg != nil {
		body.Message = msg[0]
	}
	resp.body = body
	return resp
}
