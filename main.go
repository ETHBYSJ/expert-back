package main

import (
	"expert-back/initialize"
	"expert-back/pkg/conf"
	"expert-back/router"
	"expert-back/util"
)

func init() {
	initialize.Init("./config/app.yaml")
}

func main() {
	api := router.InitRouter()
	util.Log().Info("开始监听 %s", conf.SystemConfig.System.Listen)
	if err := api.Run(conf.SystemConfig.System.Listen); err != nil {
		util.Log().Error("无法监听[%s], %s", conf.SystemConfig.System.Listen, err)
	}
}
