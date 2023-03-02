package mdCos

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"happy-paradise-golang/utils"
	"time"
)

type Cdn struct {
	CdnDomain string
	CdnPKey   string
}

// 鉴权获取cdn源
// @param key string "资源路径"
// @return cUrl string "资源加密后的路径"

func (mdCos *MdCos) ObjectCdnUrl(key string) (cUrl string) {
	ghm := mdCos.generateCdnMd5Hash(key)
	cUrl = fmt.Sprintf("https://%s/%s/%s/%s", mdCos.CdnDomain, ghm.Md5Hash, ghm.HexTs, key)

	return cUrl
}

// md5hash签名
// @param ts string "强转成16进制的unix时间戳"
// @param key string "资源路径"
// @return sign string "md5hash签名串"

func (mdCos *MdCos) Sign(ts, key string) string {
	uri := fmt.Sprintf("%s/%s%s", mdCos.CdnPKey, key, ts)
	signUri := []byte(uri)
	hash := md5.New()
	hash.Write(signUri)

	return hex.EncodeToString(hash.Sum(nil))
}

// 鉴权签名
// @param key string "资源路径"
// @return gh GenerateHashMap

func (mdCos *MdCos) generateCdnMd5Hash(key string) GenerateHashMap {
	now := time.Now().Unix()
	ts := utils.MakeHex(int(now))

	sign := mdCos.Sign(ts, key)

	gh := GenerateHashMap{
		HexTs:   ts,
		Md5Hash: sign,
	}
	return gh
}
