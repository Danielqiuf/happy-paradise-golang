package global

import "time"

const (
	PrdName                                = "md-media"
	ProEnvName                             = "md-media-pro"
	ProProdName                            = "PROD"
	MiddlewareCorsOrigin                   = "*"
	MiddlewareCorsHeaders                  = "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,X-Sign-Token,api-idempotence-token"
	MiddlewareCorsMethods                  = "OPTIONS,GET,POST,PUT,PATCH,DELETE"
	MiddlewareCorsExpose                   = "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type"
	MiddlewareCorsCredentials              = "true"
	MiddlewareTransactionTxCtxKey          = "tx"
	MiddlewareTransactionForceCommitCtxKey = "ForceCommitTx"
)

const (
	Ok                  = 200
	Error               = 400
	InternalServerError = 500
)

const (
	OkMessage              = "success"
	ErrorMessage           = "failed"
	InternalServerErrorMsg = "server internal error"
)

const (
	TencentCosBucket    = "md-1304341200"
	TencentCosRegion    = "ap-guangzhou"
	TencentCosSecretId  = "AKIDZ0GLBl7F8l9ICtKXGl11nvoVzGH1SilZ"
	TencentCosSecretKey = "bqsfDWDiyxjJgcLp56sOgjYMyOaXGo7B"
	TencentCosTimeout   = 100 * time.Second
	TencentCdnBucket    = "cdn-md"
	TencentCdnAuthPkey  = "dimtm5evg50ijsx2hvuwyfoiu65"
)

const (
	TencentCdnAuthMasterSecretKey = "HOBuOSgeMh7N8"
	TencentCdnAuthSlaveSecretKey  = "465bReFFfIkgre58Wc3K4"
)

var ErrorHandle = map[int]string{
	Ok:                  OkMessage,
	Error:               ErrorMessage,
	InternalServerError: InternalServerErrorMsg,
}
