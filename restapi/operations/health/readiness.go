// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadinessHandlerFunc turns a function with the right signature into a readiness handler
type ReadinessHandlerFunc func(ReadinessParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadinessHandlerFunc) Handle(params ReadinessParams) middleware.Responder {
	return fn(params)
}

// ReadinessHandler interface for that can handle valid readiness params
type ReadinessHandler interface {
	Handle(ReadinessParams) middleware.Responder
}

// NewReadiness creates a new http.Handler for the readiness operation
func NewReadiness(ctx *middleware.Context, handler ReadinessHandler) *Readiness {
	return &Readiness{Context: ctx, Handler: handler}
}

/*Readiness swagger:route GET /readiness Health readiness

Readiness readiness API

*/
type Readiness struct {
	Context *middleware.Context
	Handler ReadinessHandler
}

func (o *Readiness) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadinessParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
