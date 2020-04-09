package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type JsonController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "liang.liangqinghai.cn"
	c.Data["Email"] = "liangqinghai@live.com"
	c.TplName = "index.tpl"
}


func (jc *JsonController) Get() {
	ctx := jc.Ctx
	responseHeader := ctx.ResponseWriter.Header();
	responseHeader.Set("Content-Type", "application/json")
	ctx.WriteString("{\"name\": \"hello\"}")
}