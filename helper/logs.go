package helper

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"os"
)

//日志变量
var Logger = logs.NewLogger()

//日志初始化
func InitLogs() {
	var logpath = GetConfig("logs", "path") //创建日志目录
	if _, err := os.Stat(logpath); err != nil {
		os.Mkdir(logpath, os.ModePerm)
	}
	var level = GetConfigInt64("logs", "level")
	if Debug {
		level = 4
	}
	maxLines := GetConfigInt64("logs", "max_lines")
	if maxLines <= 0 {
		maxLines = 10000
	}
	maxDays := GetConfigInt64("logs", "max_days")
	if maxDays <= 0 {
		maxDays = 7
	}
	//初始化日志各种配置
	LogsConf := fmt.Sprintf(`{"filename":"logs/%v.log","level":%v,"maxlines":%v,"maxsize":0,"daily":true,"maxdays":%v}`,
		GetConfig("logs", "name"), level, maxLines, maxDays)
	Logger.SetLogger(logs.AdapterFile, LogsConf)
	if Debug {
		Logger.SetLogger("console")
		Logger.Info("日志配置信息：%v", LogsConf)
	} else {
		//是否异步输出日志
		Logger.Async(1e3)
	}
	Logger.EnableFuncCallDepth(true) //是否显示文件和行号
}
