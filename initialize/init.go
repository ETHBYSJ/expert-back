package initialize

import (
	"expert-back/db"
	"expert-back/pkg/conf"
	"expert-back/service"
)

func Init(confPath string) {
	conf.Init(confPath)
	db.Init(conf.SystemConfig.Database.Connection, conf.SystemConfig.Database.Name)
	service.InitDownloadMap()
}
