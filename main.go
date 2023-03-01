package main

import (
	"embed"
	"github.com/piupuer/go-helper/pkg/listen"
	"github.com/piupuer/go-helper/pkg/log"
	"github.com/piupuer/go-helper/pkg/tracing"
	"github.com/pkg/errors"
	"happy-paradise-golang/pkg/global"
	"happy-paradise-golang/pkg/initialize"
	"happy-paradise-golang/router"
	"runtime"
	"runtime/debug"
	"strings"
)

var ctx = tracing.NewId(nil)
var conf embed.FS

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.WithContext(ctx).WithError(errors.Errorf("%v", err)).Error("[%s]project run failed, stack: %s", global.PrdName, string(debug.Stack()))
		}
	}()

	_, file, _, _ := runtime.Caller(0)
	global.RuntimeRoot = strings.TrimSuffix(file, "main.go")

	initialize.Config(ctx, conf)
	initialize.Mysql()

	listen.Http(
		listen.WithHttpCtx(ctx),
		listen.WithHttpProName(global.PrdName),
		listen.WithHttpPort(global.Conf.System.Port),
		listen.WithHttpHandler(router.RegisterServer(ctx)),
		listen.WithHttpExit(func() {
			global.Tracer.Shutdown(ctx)
		}),
	)
}
