package global

import (
	"github.com/go-sql-driver/mysql"
	_ "github.com/mitchellh/mapstructure"
	"github.com/piupuer/go-helper/pkg/log"
)

type Configuration struct {
	System SystemConfiguration `mapstructure:"system" json:"system"`
	Mysql  MysqlConfiguration  `mapstructure:"mysql" json:"mysql"`
	Redis  RedisConfiguration  `mapstructure:"redis" json:"redis"`
	Logs   LogsConfiguration   `mapstructure:"logs" json:"logs"`
}

type SystemConfiguration struct {
	Base           string `mapstructure:"-" json:"-"`
	UrlPrefix      string `mapstructure:"url-prefix" json:"urlPrefix"`
	ApiVersion     string `mapstructure:"api-version" json:"apiVersion"`
	Port           int    `mapstructure:"port" json:"port"`
	ConnectTimeout int    `mapstructure:"connect-timeout" json:"connectTimeout"`
	RateLimitMax   int64  `mapstructure:"rate-limit-max" json:"rateLimitMax"`
}

type MysqlConfiguration struct {
	Uri         string       `mapstructure:"uri" json:"uri"`
	TablePrefix string       `mapstructure:"table-prefix" json:"tablePrefix"`
	NoSql       bool         `mapstructure:"no-sql" json:"noSql"`
	Transaction bool         `mapstructure:"transaction" json:"transaction"`
	InitData    bool         `mapstructure:"initialize-data" json:"initData"`
	DSN         mysql.Config `json:"-"`
}

type RedisConfiguration struct {
	Uri          string `mapstructure:"uri" json:"uri"`
	BinlogPos    string `mapstructure:"binlog-pos" json:"binlogPos"`
	Enable       bool   `mapstructure:"enable" json:"enable"`
	EnableBinlog bool   `mapstructure:"enable-binlog" json:"enableBinlog"`
}

type LogsConfiguration struct {
	Category                 string                   `mapstructure:"category" json:"category"`
	Level                    log.Level                `mapstructure:"level" json:"level"`
	Json                     bool                     `mapstructure:"json" json:"json"`
	LineNum                  LogsLineNumConfiguration `mapstructure:"line-num" json:"lineNum"`
	OperationKey             string                   `mapstructure:"operation-key" json:"operationKey"`
	OperationAllowedToDelete bool                     `mapstructure:"operation-allowed-to-delete" json:"operationAllowedToDelete"`
}

type LogsLineNumConfiguration struct {
	Disable bool `mapstructure:"disable" json:"disable"`
	Level   int  `mapstructure:"level" json:"level"`
	Version bool `mapstructure:"version" json:"version"`
	Source  bool `mapstructure:"source" json:"source"`
}
