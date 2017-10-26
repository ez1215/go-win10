package routers

import (
	"go-win10/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/user/login", &controllers.UserController{},"post:Login")
	beego.Router("/qq", &controllers.MainController{},"get:QQLogin")
	beego.Router("/user/userCenter", &controllers.UserController{},"*:UserCenter")
}
