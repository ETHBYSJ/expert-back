package initialize

import (
	"expert-back/db"
	"expert-back/pkg/conf"
)

func Init(confPath string) {
	conf.Init(confPath)
	db.Init(conf.SystemConfig.Database.Connection, conf.SystemConfig.Database.Name)
}
