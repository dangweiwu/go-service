package config

import (
	"github.com/dangweiwu/microkit/db/mysqlx"
	"github.com/dangweiwu/microkit/db/redisx"
	"github.com/dangweiwu/microkit/observe/logx"
	"go-service/internal/apiservice/apiconfig"
)

type Config struct {
	Root  string              `yaml:"root" default:"abc" validate:"oneof=abc ./"`
	Log   logx.Config         `yaml:"log"`
	Mysql mysqlx.Config       `yaml:"mysql"`
	Redis redisx.Config       `yaml:"redis"`
	Api   apiconfig.ApiConfig `yaml:"api"`
}
