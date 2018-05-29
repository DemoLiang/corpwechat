package main

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Post()
}

func (this *MainController) Post() {
	this.Ctx.WriteString("hello world")
	return
}

