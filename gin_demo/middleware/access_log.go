package middleware

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type accessWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w accessWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func AccessLog(options ...func(*AccessLogOptions)) gin.HandlerFunc {
	ops := getAccessLogOptionsOrSetDefault(nil)
	for _, f := range options {
		f(ops)
	}
	return func(c *gin.Context) {
		startTime := time.Now()

		w := &accessWriter{
			body:           bytes.NewBuffer(nil),
			ResponseWriter: c.Writer,
		}
		c.Writer = w

		getBody(c)

		c.Next()

		// ctx := c.Request.Context()
		// ctx := tracing.RealCtx(c)
		// _, span := tracer.Start(ctx, tracing.Name(tracing.Middleware, "AccessLog"))
		// defer span.End()

		endTime := time.Now()

		// calc request exec time
		execTime := endTime.Sub(startTime).String()

		reqMethod := c.Request.Method
		reqPath := c.Request.URL.Path
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		detail := make(map[string]interface{})
		if ops.detail {
			detail = getRequestDetail(c)
			// span.SetAttributes(
			// 	attribute.String(constant.MiddlewareParamsRespLogKey, detail[constant.MiddlewareParamsRespLogKey].(string)),
			// )
		}

		detail[MiddlewareAccessLogIpLogKey] = clientIP

		// l := logger.Builder().WithFieldMap(detail)
		// 将日志输出到标准输出
		l := log.Default()

		if reqMethod == "OPTIONS" || reqPath == fmt.Sprintf("/%s/ping", ops.urlPrefix) {
			l.Printf("method: %s, path: %s, status_code: %d, exec_time: %s", reqMethod, reqPath, statusCode, execTime)

			// l.WithField("method", reqMethod).
			// 	WithField("path", reqPath).
			// 	WithField("status_code", statusCode).
			// 	WithField("exec_time", execTime).
			// 	WithContext(ctx).
			// 	Debug("AccessLog")
		} else {
			l.Printf("method: %s, path: %s, status_code: %d, exec_time: %s", reqMethod, reqPath, statusCode, execTime)

			// l.WithField("method", reqMethod).
			// 	WithField("path", reqPath).
			// 	WithField("status_code", statusCode).
			// 	WithField("exec_time", execTime).
			// 	WithContext(ctx).
			// 	Info("AccessLog")
		}
	}
}
