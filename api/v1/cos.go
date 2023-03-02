package v1

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
	"happy-paradise-golang/pkg/global"
	"happy-paradise-golang/pkg/response"
	"net/http"
	"net/url"
)

type GenerateHashMap struct {
	HexTs   string
	Md5Hash string
}

func AuthHandler(c *gin.Context) {
	authUrl := fmt.Sprintf("https://%s.cos.%s.myqcloud.com", global.TencentCosBucket, global.TencentCosRegion)
	u, _ := url.Parse(authUrl)
	base := &cos.BaseURL{
		BucketURL: u,
	}
	client := cos.NewClient(base, &http.Client{
		Timeout: global.TencentCosTimeout,
		Transport: &cos.AuthorizationTransport{
			SecretID:  global.TencentCosSecretId,
			SecretKey: global.TencentCosSecretKey,
		},
	})

	key := "media/streams/test-video/stream_master.m3u8"
	oUrl := client.Object.GetObjectURL(key)

	response.SuccessData(oUrl)
}

func generateCdnMd5Hash(key string) *GenerateHashMap {
	timestamp := "5e577978"
	uri := fmt.Sprintf("%s/%s%s", global.TencentCdnAuthPkey, key, timestamp)
	signUri := []byte(uri)
	hash := md5.New()
	hash.Write(signUri)
	fmt.Println("timestamp22", timestamp)
	gh := &GenerateHashMap{
		HexTs:   timestamp,
		Md5Hash: hex.EncodeToString(hash.Sum(nil)),
	}
	return gh
}

func AuthCdnHandler(c *gin.Context) {
	key := "stream_master.m3u8"
	ghm := generateCdnMd5Hash(key)
	cdnUrl := fmt.Sprintf("https://cdn-md.haotianhuyu.com/%s/%s/%s", ghm.Md5Hash, ghm.HexTs, key)
	response.SuccessData(cdnUrl)
}
