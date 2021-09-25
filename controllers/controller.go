package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type RegisterController struct {
	beego.Controller
}

type RegisterUser struct {
	Name     string
	Password string
}

func (this *RegisterController) Get() {
	this.TplName = "register.html"
}

func (this *RegisterController) Post() {
	// 接受返回数据
	name := this.GetString("Name")
	passwpord := this.GetString("Password")
	// 查表

	// 返回数据
	this.Ctx.WriteString("asds")
}
