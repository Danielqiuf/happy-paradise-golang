package v1

import (
	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	data := map[string]interface{}{
		"message": "uploader",
		"success": true,
	}
	c.JSON(200, data)
}
