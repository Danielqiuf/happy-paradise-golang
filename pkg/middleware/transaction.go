package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/piupuer/go-helper/pkg/tracing"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	"happy-paradise-golang/pkg/global"
	"happy-paradise-golang/pkg/response"
	"net/http"
)

type TransactionOptions struct {
	dbNoTx               *gorm.DB
	forceTransactionPath []string
}

func getTransactionOptionsOrSetDefault(options *TransactionOptions) *TransactionOptions {
	if options == nil {
		return &TransactionOptions{
			forceTransactionPath: []string{},
		}
	}
	return options
}

func Transaction(options ...func(*TransactionOptions)) gin.HandlerFunc {
	ops := getTransactionOptionsOrSetDefault(nil)
	for _, f := range options {
		f(ops)
	}
	//if ops.dbNoTx == nil {
	//	panic("dbNoTx is empty")
	//}
	return func(c *gin.Context) {
		method := c.Request.Method
		noTransaction := false
		if method == "GET" {
			noTransaction = true
		}

		defer func() {
			ctx := tracing.RealCtx(c)
			_, span := otel.Tracer(tracing.Middleware).Start(ctx, tracing.Name(tracing.Middleware, "Transaction"))
			defer span.End()

			//tx := getTx(c, *ops)
			if err := recover(); err != nil {
				if rp, ok := err.(response.Response); ok {
					//if !noTransaction {
					//	if rp.Code == global.Ok || c.GetBool(global.MiddlewareTransactionForceCommitCtxKey) {
					//		tx.Commit()
					//	} else {
					//		tx.Rollback()
					//	}
					//}
					rp.RequestId, _, _ = tracing.GetId(c)
					c.JSON(http.StatusOK, rp)
					c.Abort()
					return
				}
				//if !noTransaction {
				//	tx.Rollback()
				//}
				// 抛出异常
				panic(err)
			} else {
				//if !noTransaction {
				//	tx.Commit()
				//}
			}
			c.Abort()
		}()
		if !noTransaction {
			tx := ops.dbNoTx.Begin()
			c.Set(global.MiddlewareTransactionTxCtxKey, tx)
		}
		c.Next()
	}
}

func getTx(c *gin.Context, ops TransactionOptions) *gorm.DB {
	tx := ops.dbNoTx
	txKey, exists := c.Get(global.MiddlewareTransactionTxCtxKey)
	if exists {
		if item, ok := txKey.(*gorm.DB); ok {
			tx = item
		}
	}
	return tx
}

func WithTransactionDbTx(db *gorm.DB) func(options *TransactionOptions) {
	return func(options *TransactionOptions) {
		if db != nil {
			getTransactionOptionsOrSetDefault(options).dbNoTx = db
		}
	}
}
