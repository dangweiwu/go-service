package app

import (
	"go-service/internal/bootstrap/appctx"
	"go-service/internal/service/apiservice/app/admin/admin/adminmodel"
	"go-service/internal/service/apiservice/app/admin/auth/authmodel"
	"go-service/internal/service/apiservice/app/admin/role/rolemodel"
)

var Tables = []interface{}{
	&adminmodel.AdminPo{},
	&authmodel.AuthPo{},
	&rolemodel.RolePo{},
}

func Regdb(appctx *appctx.AppCtx) error {
	return appctx.Db.Set("gorm:ble_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(Tables...)
}
