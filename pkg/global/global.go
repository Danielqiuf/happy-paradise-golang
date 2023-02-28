package global

import (
	"context"
	"embed"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/piupuer/go-helper/pkg/log"
	"go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
)

type ConfBoxMs struct {
	Ctx context.Context
	Fs  embed.FS
	Dir string
}

func (c ConfBoxMs) Get(filename string) (bs []byte) {
	if filename == "" {
		return
	}
	f := fmt.Sprintf("%s%s%s", c.Dir, string(os.PathSeparator), filename)
	var err error
	// read from system
	bs, err = ioutil.ReadFile(f)
	if err != nil {
		log.WithContext(c.Ctx).WithError(err).Warn("[conf box]read file %s from system failed", f)
		err = nil
	}
	if len(bs) == 0 {
		// read from embed
		bs, err = c.Fs.ReadFile(f)
		if err != nil {
			log.WithContext(c.Ctx).WithError(err).Warn("[conf box]read file %s from embed failed", f)
		}
	}
	return
}

var (
	Mode        string
	Redis       redis.UniversalClient
	Mysql       *gorm.DB
	Conf        Configuration
	Tracer      *trace.TracerProvider
	RuntimeRoot string
	ConfBox     ConfBoxMs
)
