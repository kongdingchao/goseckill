package helper

import (
	"github.com/astaxie/beego"
)

//常量
const (
	//DocHub Version
	VERSION = "v2.4"
)

//全局变量
var (
	Debug = beego.AppConfig.String("runmode") == beego.DEV
)

//类型定义
type ConfigCate string
