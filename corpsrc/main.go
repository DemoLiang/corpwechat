package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"sync"
)

func enableHTTPS() {
	beego.BConfig.Listen.HTTPSCertFile = "cert/server.crt"
	beego.BConfig.Listen.HTTPSKeyFile = "cert/server.key"
	beego.BConfig.Listen.HTTPSPort = 443
	beego.BConfig.Listen.HTTPPort = 80
	beego.BConfig.Listen.EnableHTTPS = true
	beego.BConfig.Listen.EnableHTTP = true
	beego.BConfig.CopyRequestBody = true
}

func InitBeego() {
	enableHTTPS()
}

var weChatLock *sync.RWMutex
var AccessTokenMap = map[string]string{}

func main() {
	InitBeego()
	beego.SetStaticPath("/", "../www")
	beego.Router("/wechat/contact", &ContactAppController{AgentId: "contact"})
	beego.Router("/wechat/contact/user/simplelist", &ContactAppUserSimplelistController{AgentId: "contact"})
	beego.Router("/wechat/demo", &DemoAppController{AgentId: "1000002"})
	beego.Router("/wechat/demo/login", &DemoAppLoginController{AgentId: "1000002"})
	beego.Router("/*", &MainController{})

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	InitWeChatMap()
	go UpdateWeChatMap()
	beego.Run()
}
