package app

import (
	"go-service/internal/apiservice/app/admin/admin/adminmodel"
	"go-service/internal/apiservice/app/admin/auth/authmodel"
	"go-service/internal/apiservice/app/admin/role/rolemodel"
	"go-service/internal/bootstrap/appctx"
)

var Tables = []interface{}{
	&adminmodel.AdminPo{},
	&authmodel.AuthPo{},
	&rolemodel.RolePo{},
}

func Regdb(appctx *appctx.AppCtx) error {
	return appctx.Db.Set("gorm:ble_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(Tables...)
}
