package middleware

import (
	"context"
	"reflect"
	"strings"
)

const ()

type CorsOptions struct {
	origin     string
	header     string
	expose     string
	method     string
	credential string
}

func WithCorsOrigin(s string) func(*CorsOptions) {
	return func(options *CorsOptions) {
		getCorsOptionsOrSetDefault(options).origin = s
	}
}

func WithCorsHeader(s string) func(*CorsOptions) {
	return func(options *CorsOptions) {
		getCorsOptionsOrSetDefault(options).header = s
	}
}

func WithCorsExpose(s string) func(*CorsOptions) {
	return func(options *CorsOptions) {
		getCorsOptionsOrSetDefault(options).expose = s
	}
}

func WithCorsMethod(s string) func(*CorsOptions) {
	return func(options *CorsOptions) {
		getCorsOptionsOrSetDefault(options).method = s
	}
}

func WithCorsCredential(s string) func(*CorsOptions) {
	return func(options *CorsOptions) {
		getCorsOptionsOrSetDefault(options).credential = s
	}
}

func getCorsOptionsOrSetDefault(options *CorsOptions) *CorsOptions {
	if options == nil {
		return &CorsOptions{
			origin:     MiddlewareCorsOrigin,
			header:     MiddlewareCorsHeaders,
			expose:     MiddlewareCorsExpose,
			method:     MiddlewareCorsMethods,
			credential: MiddlewareCorsCredentials,
		}
	}
	return options
}

type ParamsOptions struct {
	bodyCtxKey  string
	queryCtxKey string
}

func WithParamsBodyCtxKey(key string) func(*ParamsOptions) {
	return func(options *ParamsOptions) {
		getParamsOptionsOrSetDefault(options).bodyCtxKey = key
	}
}

func WithParamsQueryCtxKey(key string) func(*ParamsOptions) {
	return func(options *ParamsOptions) {
		getParamsOptionsOrSetDefault(options).queryCtxKey = key
	}
}

func getParamsOptionsOrSetDefault(options *ParamsOptions) *ParamsOptions {
	if options == nil {
		return &ParamsOptions{
			bodyCtxKey:  MiddlewareParamsBodyCtxKey,
			queryCtxKey: MiddlewareParamsQueryCtxKey,
		}
	}
	return options
}

type AccessLogOptions struct {
	urlPrefix string
	detail    bool
}

func WithAccessLogUrlPrefix(prefix string) func(*AccessLogOptions) {
	return func(options *AccessLogOptions) {
		getAccessLogOptionsOrSetDefault(options).urlPrefix = strings.Trim(prefix, "/")
	}
}

func WithAccessLogDetail(flag bool) func(*AccessLogOptions) {
	return func(options *AccessLogOptions) {
		getAccessLogOptionsOrSetDefault(options).detail = flag
	}
}

func getAccessLogOptionsOrSetDefault(options *AccessLogOptions) *AccessLogOptions {
	if options == nil {
		return &AccessLogOptions{
			urlPrefix: MiddlewareUrlPrefix,
			detail:    true,
		}
	}
	return options
}

type PrintRouterOptions struct {
	ctx context.Context
}

func WithPrintRouterCtx(ctx context.Context) func(*PrintRouterOptions) {
	return func(options *PrintRouterOptions) {
		if !InterfaceIsNil(ctx) {
			getPrintRouterOptionsOrSetDefault(options).ctx = ctx
		}
	}
}

func getPrintRouterOptionsOrSetDefault(options *PrintRouterOptions) *PrintRouterOptions {
	if options == nil {
		return &PrintRouterOptions{
			ctx: context.Background(),
		}
	}
	return options
}

func InterfaceIsNil(i interface{}) bool {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		return v.IsNil()
	}
	return i == nil
}
