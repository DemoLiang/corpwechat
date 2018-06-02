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


* 其中type中定义了大部分的消息推送的结构定义可以通过解析下发的消息到结构
以及响应消息结构定义，xml编码响应消息返回给微信
```
//基础响应信息，返回格式
type BasicResponseInfo struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
}

//文本请求body content 内容
type TextRequestMessage struct {
	BasicRequestInfo
	Content string `xml:"Content"`
	MsgId   string `xml:"MsgId"`
}

//文本请求响应body content 内容
type TextResponseMessage struct {
	BasicResponseInfo
	Content string `xml:"Content"`
}

//图片消息请求body content 内容
type ImageRequestMessage struct {
	BasicRequestInfo
	PicUrl   string `xml:"PicUrl"`
	MedialId string `xml:"MedialId"`
}

//图片消息回复消息格式
type ImageResponseMessage struct {
	BasicResponseInfo
	Image struct {
		MediaId string `xml:"MediaId"`
	} `xml:"Image"`
}
```