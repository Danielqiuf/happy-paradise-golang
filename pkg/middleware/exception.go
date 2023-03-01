package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/piupuer/go-helper/pkg/log"
	"github.com/piupuer/go-helper/pkg/tracing"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"happy-paradise-golang/pkg/global"
	"happy-paradise-golang/pkg/response"
	"net/http"
	"runtime/debug"
)

func Exception(c *gin.Context) {
	defer func() {
		ctx := tracing.RealCtx(c)
		_, span := otel.Tracer(tracing.Middleware).Start(ctx, tracing.Name(tracing.Middleware, "Exception"))
		defer span.End()
		if err := recover(); err != nil {
			e := errors.Errorf("%v", err)
			log.WithContext(c).WithError(e).Error("runtime exception, stack: %s", string(debug.Stack()))
			rp := response.Response{
				Code: global.InternalServerError,
				Data: map[string]interface{}{},
				Msg:  global.ErrorHandle[global.InternalServerError],
			}
			rp.RequestId, _, _ = tracing.GetId(c)
			span.RecordError(e)
			// set json data
			c.JSON(http.StatusOK, rp)
			c.Abort()
			return
		}
	}()
	c.Next()
}
