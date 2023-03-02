package v1

import (
	"github.com/gin-gonic/gin"
	"happy-paradise-golang/pkg/response"
)

func StreamVideoHandler(c *gin.Context) {
	rp := response.StreamVideoResponse{
		VideoName: "测试视频test1",
		Source:    "https://md-1304341200.cos.ap-guangzhou.myqcloud.com/media/streams/test-video/video1.m3u8",
	}

	rp.Id = 12
	response.SuccessData(rp)
}

func StreamVideoUploadHandler(c *gin.Context) {

}
