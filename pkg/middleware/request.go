package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/piupuer/go-helper/pkg/constant"
	"github.com/piupuer/go-helper/pkg/tracing"
	"io/ioutil"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", constant.MiddlewareCorsOrigin)
		c.Header("Access-Control-Allow-Headers", constant.MiddlewareCorsHeaders)
		c.Header("Access-Control-Allow-Methods", constant.MiddlewareCorsMethods)
		c.Header("Access-Control-Expose-Headers", constant.MiddlewareCorsExpose)
		c.Header("Access-Control-Allow-Credentials", constant.MiddlewareCorsCredentials)

		c.Next()
	}
}

func RequestId(c *gin.Context) {
	requestId, _, _ := tracing.GetId(c)
	if requestId == "" {
		c.Request = c.Request.WithContext(tracing.NewId(c))
	}
	c.Next()
}

func SecurityHeader(c *gin.Context) {
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")
	c.Header("X-Frame-Options", "deny")
	c.Next()
}

func Params(c *gin.Context) {
	getBody(c)
	getQuery(c)
	c.Next()
}

func getBody(c *gin.Context) (rp string) {
	if v := c.GetString(constant.MiddlewareParamsBodyCtxKey); v != "" {
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
		rp = constant.MiddlewareParamsNullBody
	} else {
		rp = string(body)
	}
	c.Set(constant.MiddlewareParamsBodyCtxKey, rp)
	return
}

func getQuery(c *gin.Context) (rp string) {
	if v := c.GetString(constant.MiddlewareParamsQueryCtxKey); v != "" {
		rp = v
		return
	}
	rp = c.Request.URL.RawQuery
	c.Set(constant.MiddlewareParamsQueryCtxKey, rp)
	return
}
