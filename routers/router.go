package routers

import (
	"QuickBeeGo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/aa/", &controllers.MainController{})
}
