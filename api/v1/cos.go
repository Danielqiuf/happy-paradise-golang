package v1

import (
	"github.com/gin-gonic/gin"
	"happy-paradise-golang/pkg/response"
	"happy-paradise-golang/utils/mdCos"
)

type GenerateHashMap struct {
	HexTs   string
	Md5Hash string
}

func AuthHandler(c *gin.Context) {
	mdc := mdCos.New()
	oUrl := mdc.ObjectUrl(mdCos.GetStreamVideoLocation("stream_master"))

	response.SuccessData(oUrl)
}

func AuthCdnHandler(c *gin.Context) {
	mdc := mdCos.New()
	cdnUrl := mdc.ObjectCdnUrl(mdCos.GetStreamVideoLocation("stream_master"))
	response.SuccessData(cdnUrl)
}
