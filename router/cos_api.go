package router

import (
	"github.com/gin-gonic/gin"
	v1 "happy-paradise-golang/api/v1"
)

func CosApi(group *gin.RouterGroup) {
	cosApiGroup := group.Group("/cos")

	cosApiGroup.GET("/auth", v1.AuthHandler)

	cosApiGroup.GET("/cdnauth", v1.AuthCdnHandler)
}
