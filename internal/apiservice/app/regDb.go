package app

import "go-service/internal/bootstrap/appctx"

var Tables = []interface{}{}

func Regdb(appctx *appctx.AppCtx) error {
	return appctx.Db.Set("gorm:ble_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(Tables...)
}
