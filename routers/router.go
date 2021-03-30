package routers

import (
	"QuickBeeGo/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.MainController{})
	web.Router("/t/", &controllers.MainController{}, "post:T1")
	//web.Get("/t2",controllers.InfController.T2)
	//web.Router("/t3", &controllers.InfController{},"post:T3")
}
