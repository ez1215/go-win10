package controllers

import (
	"github.com/astaxie/beego"
	"log"
	. "go-win10/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Login() {
	var username string = c.GetString("username")
	var passwrod string = c.GetString("password")

	log.Println(username)
	user, err := CheckUser(username, passwrod)

	if err != nil {
		log.Println(err)
		c.TplName = "index.tpl"
		c.Data["result"] = "用户名或密码错误"
		c.Data["status"] = true
		return
	}

	if user != nil {
		c.SetSession("user", user)
		log.Println("登录成功")
		c.Ctx.Redirect(302, "/user/userCenter")
	}
}

func (c *UserController) UserCenter() {
	c.TplName = "usercenter.tpl"
	u := c.GetSession("user")
	log.Println(u)
	c.Data["user"] = u
	return
}