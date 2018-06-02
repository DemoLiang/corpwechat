# corpwechat
企业微信扫码授权登录组件

简单的验证企业微信扫码授权登录，登录后后去到用户的UserId后即可拿用户的UserId换取企业自己服务器后台的token

- 获取代码  ` go get -v github.com/DemoLiang/corpwechat ` 

增加了基础的接口测试函数封装，可以直接使用函数，配置多个APP的处理函数然后增加自己的handler处理，解析客户端发来的信息

多个APP需要在const文件配置响应的APP ID的基础信息

```
var CorpAppMap = map[string]CorpWeChatApp{
	//"XXXX": CorpWeChatApp{
	//	EncodingAESKey: "XXXXX",
	//	Token:          "XXXX",
	//	CoprId:         "1000002",
	//	Seceret:        "XXXXX",
	//},
}
```

同时在多个APP直接，需要在接口调用路由中配置相应的AgentId
```
beego.Router("/wechat/demo", &DemoAppController{AgentId: "1000002"})
```
