package routers

import (
	"github.com/astaxie/beego"
	"golang-myprojects/goseckill/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/seckill", &controllers.SeckillController{})
}
