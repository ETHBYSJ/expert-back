package main

import (
	"expert-back/initialize"
	"expert-back/pkg/conf"
	util2 "expert-back/pkg/util"
	"expert-back/router"
)

func init() {
	initialize.Init("./config/app.yaml")
}

func main() {
	api := router.InitRouter()
	util2.Log().Info("开始监听 %s", conf.SystemConfig.System.Listen)
	if err := api.Run(conf.SystemConfig.System.Listen); err != nil {
		util2.Log().Error("无法监听[%s], %s", conf.SystemConfig.System.Listen, err)
	}
}
