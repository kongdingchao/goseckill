package main

import (
	"github.com/astaxie/beego"
	"golang-myprojects/goseckill/helper"
	_ "golang-myprojects/goseckill/routers"
)

func main() {
	beego.Run()
}

func init() {
	helper.Logger.Info("Powered By kdc")
	helper.Logger.Info("Version:%v", helper.VERSION)
}
