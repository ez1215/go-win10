package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

type GameController struct {
	beego.Controller
}

func (c *GameController) Flower()  {
	log.Println("玫瑰花")
	c.TplName = "flower.tpl"
}

func (c *GameController) Butterfly()  {
	log.Println("蝴蝶")
	c.TplName = "butterfly.tpl"
}

func (c *GameController) Jfcz()  {
	log.Println("见缝插针")
	c.TplName = "jfcz.tpl"
}

func (c *GameController) Lianliankan()  {
	log.Println("连连看")
	c.TplName = "lianliankan.tpl"
}

func (c *GameController) Sankbox()  {
	log.Println("推箱子")
	c.TplName = "sankbox.tpl"
}

func (c *GameController) Thunder()  {
	log.Println("扫雷")
	c.TplName = "thunder.tpl"
}

func (c *GameController) Tank()  {
	log.Println("坦克大战")
	c.TplName = "tank.tpl"
}

func (c *GameController) Snake()  {
	log.Println("贪吃蛇")
	c.TplName = "snake.tpl"
}