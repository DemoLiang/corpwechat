package main

import (
	"github.com/astaxie/beego"
)

type DemoAppLoginController struct {
	beego.Controller
	AgentId string
}

func (this *DemoAppLoginController) Get() {
	this.Post()
}

func (this *DemoAppLoginController) Post() {
	code := this.GetString("code")
	state := this.GetString("state")

	if state == "" || code == "" {
		this.Ctx.WriteString("login fail")
		return
	}

	userid, _ := GetkUserInfo(AccessTokenMap[this.AgentId], code)
	//TODO handler login

	this.Ctx.WriteString("login success userid:" + userid)
	return
}