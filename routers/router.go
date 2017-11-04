package routers

import (
	"go-win10/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.UserController{},"get:ToIndex")
	beego.Router("/user/login", &controllers.UserController{},"post:Login")
	beego.Router("/qq", &controllers.MainController{},"get:QQLogin")
	beego.Router("/logout", &controllers.UserController{},"post:LogOut")
	beego.Router("/index", &controllers.UserController{},"get:ToIndex")
	beego.Router("/desktop", &controllers.UserController{},"get:Desktop")

	beego.Router("/flower", &controllers.GameController{},"get:Flower")
	beego.Router("/butterfly", &controllers.GameController{},"get:Butterfly")
	beego.Router("/jfcz", &controllers.GameController{},"get:Jfcz")
	beego.Router("/lianliankan", &controllers.GameController{},"get:Lianliankan")
	beego.Router("/sankbox", &controllers.GameController{},"get:Sankbox")
	beego.Router("/thunder", &controllers.GameController{},"get:Thunder")
	beego.Router("/tank", &controllers.GameController{},"get:Tank")
	beego.Router("/snake", &controllers.GameController{},"get:Snake")

}
