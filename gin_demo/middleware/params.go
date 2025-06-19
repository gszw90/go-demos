package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	MiddlewareParamsBodyCtxKey           = "ParamsBody"
	MiddlewareParamsQueryCtxKey          = "ParamsQuery"
	MiddlewareParamsNullBody             = "{}"
	MiddlewareParamsQueryLogKey          = "Query"
	MiddlewareParamsBodyLogKey           = "Body"
	MiddlewareParamsRespCtxKey           = "ResponseBody"
	MiddlewareParamsRespLogKey           = "Resp"
	HITPAY_CALLBACK_HEADER_SIGNATURE_KEY = "Hitpay-Signature"

	MiddlewareCorsOrigin      = "*"
	MiddlewareCorsHeaders     = "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,X-Sign-Token,api-idempotence-token"
	MiddlewareCorsMethods     = "OPTIONS,GET,POST,PUT,PATCH,DELETE"
	MiddlewareCorsExpose      = "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type"
	MiddlewareCorsCredentials = "true"

	MiddlewareUrlPrefix         = "api"
	MiddlewareAccessLogIpLogKey = "Ip"
)

func Params(c *gin.Context) {

	getBody(c)
	getQuery(c)
	c.Next()
}

func getBody(c *gin.Context) (rp string) {
	if v := c.GetString(MiddlewareParamsBodyCtxKey); v != "" {
		rp = v
		return
	}
	reqMethod := c.Request.Method
	// read body
	var body []byte
	if reqMethod == http.MethodPost || reqMethod == http.MethodPut || reqMethod == http.MethodPatch {
		var err error
		body, err = ioutil.ReadAll(c.Request.Body)
		if err == nil {
			// write back to gin request body
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
	}
	if len(body) == 0 {
		rp = MiddlewareParamsNullBody
	} else {
		rp = string(body)
	}
	c.Set(MiddlewareParamsBodyCtxKey, rp)
	return
}

func getQuery(c *gin.Context) (rp string) {
	if v := c.GetString(MiddlewareParamsQueryCtxKey); v != "" {
		rp = v
		return
	}
	rp = c.Request.URL.RawQuery
	c.Set(MiddlewareParamsQueryCtxKey, rp)
	return
}

func getResp(c *gin.Context) (rp string) {
	if v := c.GetString(MiddlewareParamsRespCtxKey); v != "" {
		rp = v
		return
	}
	if w, ok := c.Writer.(*accessWriter); ok {
		rp = w.body.String()
		c.Set(MiddlewareParamsRespCtxKey, rp)
	}
	return
}

func getRequestDetail(c *gin.Context) (rp map[string]interface{}) {
	rp = make(map[string]interface{})
	ct := c.Writer.Header().Get("Content-Type")
	var d1 string
	if strings.Contains(ct, "application/json") ||
		strings.Contains(ct, "text/plain") ||
		ct == "" {
		d1 = trim(getResp(c))
	} else {
		d1 = fmt.Sprintf("`%s`", ct)
	}
	rp[MiddlewareParamsRespLogKey] = d1
	rp[MiddlewareParamsQueryLogKey] = trim(getQuery(c))
	rp[MiddlewareParamsBodyLogKey] = trim(getBody(c))
	return
}

func trim(s string) string {
	s = compact(s)
	s = strings.ReplaceAll(s, "\"", "'")
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, "\n", "")
	if len(s) > 500 {
		s = fmt.Sprintf("%s......omitted......%s", s[0:250], s[len(s)-250:len(s)])
	}
	return s
}

func compact(s string) string {
	got := new(bytes.Buffer)
	if err := json.Compact(got, []byte(s)); err != nil {
		return s
	}
	return got.String()
}
