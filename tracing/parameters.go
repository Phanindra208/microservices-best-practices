package tracing

import (
	"context"
	"fmt"

	"github.com/iAmPlus/microservice/config"
	"github.com/opentracing/opentracing-go"
)

// Request directions
const (
	DirectionIncoming = "incoming"
	DirectionOutgoing = "outgoing"
)

type spanTagType string

// Span tags
const (
	SchemaVersion spanTagType = "span_tag_schema_version"
	API           spanTagType = "span_tag_api"
	Name          spanTagType = "span_tag_name"
)

var allSpanTags = []spanTagType{
	Name,
}

// TracingContext is just a simple wrapper for the interface.
type TracingContext interface {
	With(spanTag spanTagType, value string) TracingContext
	Ctx() context.Context
	context.Context
}

type wrappedContext struct {
	context.Context
}

// Context returns a new tracing context.
func Context(ctx context.Context) TracingContext {
	return &wrappedContext{Context: ctx}
}

// With sets a tracing context value.
func (wctx *wrappedContext) With(spanTag spanTagType, value string) TracingContext {
	wctx.Context = context.WithValue(wctx.Context, spanTag, value)
	return wctx
}

// Ctx returns wrapped context.Context.
func (wctx *wrappedContext) Ctx() context.Context {
	return wctx.Context
}

// WithValue sets single context.Context value.
func WithValue(ctx context.Context, spanTag spanTagType, value string) context.Context {
	return context.WithValue(ctx, spanTag, value)
}

func configureSpan(ctx context.Context, span opentracing.Span, direction string) {
	version := ctx.Value(SchemaVersion)
	if version == nil {
		version = config.Vars.ZipkinSchemaVersion
	}
	api := ctx.Value(API)
	if api == nil {
		api = config.Vars.ZipkinServiceName + "-unknown"
	}
	name := ctx.Value(Name)
	if name == nil {
		name = "unknownoperation"
	}
	span.SetOperationName(fmt.Sprintf("%s_%s_%s", direction, api, name))
	span.SetTag("schema_version", version)
	span.SetTag("request_direction", direction)
	span.SetTag("targeted_api", api)
}
