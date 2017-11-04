package controllers

import (
	"github.com/astaxie/beego"
	"log"
	. "go-win10/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) ToIndex()  {
	log.Println("注销")
	c.TplName = "login.tpl"
}

func (c *UserController) Login() {
	var username string = c.GetString("username")
	var passwrod string = c.GetString("password")

	if(username == "" || passwrod==""){
		c.TplName = "login.tpl"
		c.Data["result"] = "用户名或密码错误"
		c.Data["status"] = true
		return
	}
	log.Println(username)
	user, err := CheckUser(username, passwrod)

	if err != nil {
		log.Println("查询用户异常",err)
		c.TplName = "login.tpl"
		c.Data["result"] = "用户名或密码错误"
		c.Data["status"] = true
		return
	}

	if user != nil {
		c.SetSession("user", user)
		log.Println("登录成功")
		c.Ctx.Redirect(302, "/desktop")
	}
}

func (c *UserController) LogOut()  {
	var u  = c.GetSession("user");
	log.Println("注销用户",u)
	c.DelSession("user")
	c.Data["json"] = "ok"
	c.ServeJSON()
}

func (c *UserController) Desktop() {
	c.TplName = "desktop.tpl"
	u := c.GetSession("user")
	log.Println(u)
	c.Data["user"] = u
	return
}