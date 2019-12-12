package controllers

import (
	"github.com/astaxie/beego"
	"golang-myprojects/goseckill/models"
)

type SeckillController struct {
	beego.Controller
}

func (c *SeckillController) Get() {
	var s = models.GetSeckill()
	c.Data["ComName"] = s.ComName
	c.Data["ComNumber"] = s.ComNumber
	c.TplName = "seckill.tpl"
}
