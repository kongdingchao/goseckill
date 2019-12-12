package helper

import "github.com/astaxie/beego"

func init() {
	//初始话配置
	InitConfig(beego.AppConfig.String("logconf"))

	//初始化日志
	InitLogs()
}
