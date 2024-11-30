package option

import (
	"github.com/dangweiwu/microkit/yamlconfig"
	"go-service/internal/bootstrap"
	"go-service/internal/config"
)

/*
*
启动api
*/
type RunService struct {
	ApiHost string `long:"apihost" description:"api启动host"`
}

func (this *RunService) Execute(args []string) error {
	var c config.Config
	yamlconfig.MustLoad(Opt.ConfigPath, &c)

	boot := bootstrap.NewBootStrap()
	err := boot.Init(c)
	if err != nil {
		return err
	}
	boot.Run()
	return nil
}
