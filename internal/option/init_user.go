package option

import (
	"context"
	"github.com/dangweiwu/microkit/yamlconfig"
	"go-service/internal/apiservice/app/admin/admin/adminmodel"
	"go-service/internal/apiservice/pkg"
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/config"
	"gorm.io/gorm"
	"log"
)

type InitSuperUser struct {
	Password string `long:"password" short:"p" description:"超级管理员设置密码"`
}

func (this *InitSuperUser) Usage() string {
	return `//设置超级管理员密码`
}

func (this *InitSuperUser) Execute(args []string) error {
	var c config.Config
	yamlconfig.MustLoad(Opt.ConfigPath, &c)

	ctx, cf := context.WithCancel(context.Background())

	ctx2, err := appctx.NewAppCtx(ctx, cf, c)
	if err != nil {

		return err
	}

	po := &adminmodel.AdminPo{}
	pw, err := pkg.GetPassword(this.Password, ctx2.Config.Salt)

	if err != nil {
		return err
	}

	r := ctx2.Db.Model(po).Where("account=?", "admin").Take(po)
	if r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			log.Println("初始化超级管理员")
			po.Account = "admin"
			po.Name = "超级管理员"

			po.Password = pw
			po.Role = "admin"
			po.Status = "1"
			r := ctx2.Db.Create(po)
			if r.Error != nil {
				log.Panic(r.Error)
			}
			return nil
		} else {
			return err
		}
	}

	if r := ctx2.Db.Model(po).Update("password", pw); r.Error != nil {
		return r.Error
	}

	return nil
}
