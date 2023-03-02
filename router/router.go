package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/piupuer/go-helper/pkg/log"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"happy-paradise-golang/pkg/global"
	"happy-paradise-golang/pkg/middleware"
)

func RegisterServer(ctx context.Context) *gin.Engine {
	r := gin.New()

	gin.DebugPrintRouteFunc = func(httpMethod string, absolutePath string, handlerName string, nuHandlers int) {
		log.WithContext(ctx).Debug("[gin-route] %-6s %-40s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	//
	r.Use(
		otelgin.Middleware(global.PrdName),
		middleware.Cors(),
		middleware.SecurityHeader,
		middleware.Params,
		middleware.RequestId,
		//middleware.Sign()
		middleware.AccessLog(
			middleware.WithAccessLogUrlPrefix(global.Conf.System.UrlPrefix),
		),
		middleware.Exception,
		middleware.Transaction(
			middleware.WithTransactionDbTx(global.Mysql),
		),
	)

	apiGroup := r.Group(global.Conf.System.UrlPrefix)

	apiGroup.GET("/ping")

	v1Group := apiGroup.Group(global.Conf.System.ApiVersion)

	StreamApi(v1Group)

	CosApi(v1Group)

	return r
}
