package mdCos

import (
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
	"happy-paradise-golang/pkg/global"
	"net/http"
	"net/url"
	"time"
)

type MdCos struct {
	CosBucket    string
	CosRegion    string
	CosSecretId  string
	CosSecretKey string
	CosTimeout   time.Duration
	Cdn
}

type GenerateHashMap struct {
	HexTs   string
	Md5Hash string
}

func New() *MdCos {
	mdCos := &MdCos{
		CosBucket:    global.TencentCosBucket,
		CosRegion:    global.TencentCosRegion,
		CosSecretId:  global.TencentCosSecretId,
		CosSecretKey: global.TencentCosSecretKey,
		CosTimeout:   global.TencentCosTimeout,
	}
	mdCos.CdnDomain = global.TencentCdnDomain
	mdCos.CdnPKey = global.TencentCdnAuthMasterSecretKey
	return mdCos
}

func (mdCos *MdCos) ObjectUrl(key string) (oUrl *url.URL) {
	authUrl := fmt.Sprintf("https://%s.cos.%s.myqcloud.com", mdCos.CosBucket, mdCos.CosRegion)
	u, _ := url.Parse(authUrl)
	base := &cos.BaseURL{
		BucketURL: u,
	}
	client := cos.NewClient(base, &http.Client{
		Timeout: mdCos.CosTimeout,
		Transport: &cos.AuthorizationTransport{
			SecretID:  mdCos.CosSecretId,
			SecretKey: mdCos.CosSecretKey,
			Transport: &debug.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   false,
			},
		},
	})

	oUrl = client.Object.GetObjectURL(key)

	return oUrl
}
