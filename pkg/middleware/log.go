package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/piupuer/go-helper/pkg/constant"
	"github.com/piupuer/go-helper/pkg/log"
	"happy-paradise-golang/pkg/global"
	"strings"
	"time"
)

type AccessLogOptions struct {
	urlPrefix string
	detail    bool
}

type accessWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func AccessLog(options ...func(*AccessLogOptions)) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		w := &accessWriter{
			body:           bytes.NewBuffer(nil),
			ResponseWriter: c.Writer,
		}
		c.Writer = w

		getBody(c)

		c.Next()

		endTime := time.Now()

		// calc request exec time
		execTime := endTime.Sub(startTime).String()

		reqMethod := c.Request.Method
		reqPath := c.Request.URL.Path
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		detail := make(map[string]interface{})

		detail[constant.MiddlewareAccessLogIpLogKey] = clientIP

		l := log.WithContext(c).WithFields(detail)

		if reqMethod == "OPTIONS" || reqPath == fmt.Sprintf("/%s/ping", global.Conf.System.UrlPrefix) {
			l.Debug(
				"%s %s %d %s",
				reqMethod,
				reqPath,
				statusCode,
				execTime,
			)
		} else {
			l.Info(
				"%s %s %d %s",
				reqMethod,
				reqPath,
				statusCode,
				execTime,
			)
		}
	}
}

func WithAccessLogUrlPrefix(prefix string) func(*AccessLogOptions) {
	return func(options *AccessLogOptions) {
		options.urlPrefix = strings.Trim(prefix, "/")
	}
}
