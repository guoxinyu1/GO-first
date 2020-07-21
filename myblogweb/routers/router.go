package routers

import (
	"github.com/astaxie/beego"
	"myblogweb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})
}
