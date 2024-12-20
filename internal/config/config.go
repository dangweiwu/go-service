package config

import (
	"github.com/dangweiwu/microkit/db/mysqlx"
	"github.com/dangweiwu/microkit/db/redisx"
	"github.com/dangweiwu/microkit/observe/logx"
	"go-service/internal/apiservice/apiconfig"
	"go-service/internal/apiservice/pkg/jwtx/jwtconfig"
	"go-service/internal/service/pprofservice/pprofcfg"
)

type Config struct {
	Root  string              `yaml:"root" default:"./" validate:"required"`
	Log   logx.Config         `yaml:"log"`
	Mysql mysqlx.Config       `yaml:"mysql"`
	Redis redisx.Config       `yaml:"redis"`
	Api   apiconfig.ApiConfig `yaml:"api"`
	Pprof pprofcfg.Config     `yaml:"pprof"`
	Jwt   jwtconfig.JwtConfig `yaml:"jwt"`
}
