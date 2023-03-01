package router

import (
	"github.com/gin-gonic/gin"
	v1Handler "happy-paradise-golang/api/v1"
)

func StreamApi(group *gin.RouterGroup) {
	group.GET("/stream/video/:id", v1Handler.StreamVideoHandler)
	group.POST("/stream/upload", v1Handler.StreamVideoHandler)
}
