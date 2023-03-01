package global

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

var ErrorHandle = map[int]string{
	Ok:                  OkMessage,
	Error:               ErrorMessage,
	InternalServerError: InternalServerErrorMsg,
}
