package main

import (
	"github.com/astaxie/beego"
)

const (
	AgentId = "1000002"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.Post()
}

func (this *LoginController) Post() {
	code := this.GetString("code")
	state := this.GetString("state")

	if state == ""|| code == ""{
		this.Ctx.WriteString("login fail")
		return
	}

	userid,_:=GetkUserInfo(AccessTokenMap[AgentId],code)
	//TODO handler login

	this.Ctx.WriteString("login success userid:"+userid)
	return
}