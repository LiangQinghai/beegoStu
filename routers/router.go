package routers

import (
	"github.com/liangqinghai/beegoStu/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/json", &controllers.JsonController{})
}
