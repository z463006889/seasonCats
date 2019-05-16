package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

//配置base controller来给所有的controller继承使用
type baseController struct {
	beego.Controller
	i18n.Locale
}
