package models

import (
	"github.com/astaxie/beego/orm"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type UserEntity struct {
	Id       int
	Name     string
	Age      int
	Birth    string
	Email    string
	Password string
	Avator   string
	Nickname string
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/bee?charset=utf8")
	orm.Debug=true
}

func CheckUser(loginname string, password string) (u *UserEntity, e error) {
	log.Println("用户:", loginname, ",密码:", password)
	orm.Debug = true
	o := orm.NewOrm()

	var user UserEntity
	err := o.Raw("select * from user where name = ? and password = ?", loginname, password).QueryRow(&user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, err
}
