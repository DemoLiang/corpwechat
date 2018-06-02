package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"sort"
	"strconv"
	"strings"
	"time"
)

type DemoAppController struct {
	beego.Controller
	AgentId string
}

func (this *DemoAppController) Get() {
	msgSignature := this.GetString("msg_signature")
	signature := this.GetString("signature")
	timestamp := this.GetString("timestamp")
	nonce := this.GetString("nonce")
	echostr := this.GetString("echostr")
	encryptType := this.GetString("encrypt_type")

	if msgSignature == "" || echostr == "" || nonce == "" || timestamp == "" {
		Log("url check fail\n")
		return
	}
	if encryptType == "raw" {
		//配置明文请求URL 的校验
		var list = []string{CorpAppMap[this.AgentId].Token, timestamp, nonce}
		sort.Strings(list)
		sigStr := strings.Join(list, "")

		sha1 := sha1.New()
		sha1.Write([]byte(sigStr))
		hash := sha1.Sum(nil)
		Log("signature:%v \n timestamp:%v \n nonce:%v\n echostr:%v \n sigStr:%v \n list:%v\n hash:%v\n", signature, timestamp, nonce, echostr, sigStr, list, fmt.Sprintf("%x", hash))
		if signature == fmt.Sprintf("%x", hash) {
			Log("hash eq echostr ec")
			this.Ctx.Output.Body([]byte(echostr))
		}
		return
	} else {
		//配置加密请求URL 的校验
		aesMsg, _ := base64.StdEncoding.DecodeString(echostr)
		orgMsg, err := AesDecrypt(aesMsg, []byte(CorpAppMap[this.AgentId].AesKey))
		if err != nil {
			panic(err)
		}
		msg, _ := ParseEncryptRequestBodyContent(orgMsg)
		this.Ctx.WriteString(string(msg))
		return
	}
}

func (this *DemoAppController) Post() {
	timestamp := this.GetString("timestamp")
	nonce := this.GetString("nonce")
	msgSignatureIn := this.GetString("msg_signature")

	//获取HTTP请求的body内容
	encryptRequestBody := ParseEncryptRequestBody(this.Ctx.Input.RequestBody)
	//校验signature
	if !ValidateMsg(CorpAppMap[this.AgentId].Token, timestamp, nonce, encryptRequestBody.Encrypt, msgSignatureIn) {
		Log("invalid request\n")
		return
	}
	//解密
	aesMsg, _ := base64.StdEncoding.DecodeString(encryptRequestBody.Encrypt)
	orgMsg, err := AesDecrypt(aesMsg, []byte(CorpAppMap[this.AgentId].AesKey))
	if err != nil {
		panic(err)
	}
	//处理微信的请求数据
	rspInfo := this.WechatHandler(orgMsg)
	if rspInfo == "" {
		this.Ctx.WriteString(string(""))
		return
	}
	//响应数据
	this.Ctx.WriteString(string(rspInfo))
	return
}

//微信请求处理
func (this *DemoAppController) WechatHandler(orgMsg []byte) string {
	var basicReqInfo *BasicRequestInfo
	var rspInfo string

	var respNonce string
	var respTimestamp string

	//parse 得到请求的content 内容
	BodyContent, _ := ParseEncryptRequestBodyContent(orgMsg)
	//XML解码到基础的请求结构
	basicReqInfo, _ = ParseEncryptBasicRequestInfo(BodyContent)

	//case 所有的类型，做响应的handler 处理
	switch basicReqInfo.MsgType {
	case REQ__MSG_TYPE__TEXT:
		textReqMsg, _ := ParseEncryptTextRequestBody(BodyContent)
		Log("%v->%v\n", textReqMsg.FromUserName, textReqMsg.Content)
		rspInfo, _ = TextMessageHandler(textReqMsg)
	default:
		Log("msgtype is :%v\n", basicReqInfo.MsgType)
		rspInfo = "success"
	}
	Log("rspInfo:%v\n", rspInfo)
	respNonce = GenRandomString(8)
	respTimestamp = strconv.Itoa(int(time.Now().Unix()))
	//生成响应的response body内容，直接返回次内容
	resp, _ := MakeEncryptResponseBody(CorpAppMap[this.AgentId].Token, rspInfo, respNonce, respTimestamp, CorpAppMap[this.AgentId].AesKey)
	return string(resp)
}

//文本信息处理demo(原样返回数据)
func TextMessageHandler(textReqMsg *TextRequestMessage) (content string, err error) {
	var textRespMsg TextResponseMessage
	textRespMsg.FromUserName = textReqMsg.ToUserName
	textRespMsg.ToUserName = textReqMsg.FromUserName
	textRespMsg.MsgType = textReqMsg.MsgType
	textRespMsg.CreateTime = time.Now().Unix()
	textRespMsg.Content = textReqMsg.Content
	cnt, err := xml.Marshal(textRespMsg)
	if err != nil {
		Log("XML Marshal error:%v\n", err)
		return "", err
	}
	return string(cnt), err
}
