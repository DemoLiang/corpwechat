package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ContactAppController struct {
	beego.Controller
	AgentId string
}

func (this *ContactAppController) Get() {
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

		// Read length
		buf := bytes.NewBuffer(orgMsg[16:20])
		var length int32
		binary.Read(buf, binary.BigEndian, &length)

		msg := orgMsg[20 : 20+length]
		this.Ctx.WriteString(string(msg))
		return
	}
}

func (this *ContactAppController) Post() {
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
	Log("encryptRequestBody:%v CorpAppMap[this.AgentId].AesKey:%v\n", encryptRequestBody, CorpAppMap[this.AgentId].AesKey)
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

func (this *ContactAppController) WechatHandler(orgMsg []byte) string {
	var basicReqInfo *BasicRequestInfo
	var rspInfo string

	var respNonce string
	var respTimestamp string

	//parse 得到请求的content 内容
	BodyContent, _ := ParseEncryptRequestBodyContent(orgMsg)
	//XML解码到基础的请求结构
	basicReqInfo, _ = ParseEncryptBasicRequestInfo(BodyContent)

	//case 所有的类型，做响应的handler 处理
	switch basicReqInfo.Event {
	case REQ__EVENT_TYPE__CHANGE_CONTACT:
		changeContactReqMsg, _ := ParseEncryptEventChangeContact(BodyContent)

		if changeContactReqMsg.ChangeType == REQ__EVENT_TYPE__MSGTYPE__CREATE_USER {
			Log("REQ__EVENT_TYPE__MSGTYPE__CREATE_USER:%v->%v\n", changeContactReqMsg.FromUserName, changeContactReqMsg)
			//TODO 解析完后，同步数据到自己的数据库，更新自己的系统数据
		}
		rspInfo = "success"
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

//获取部门成员
type ContactAppUserSimplelistController struct {
	beego.Controller
	AgentId string
}

func (this *ContactAppUserSimplelistController) Get() {
	departmentId := this.GetString("department_id")
	fetchChild := this.GetString("fetch_child")
	ReqUrl := "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist"
	ReqQueryHead := map[string]string{
		"department_id": departmentId,
		"access_token":  AccessTokenMap[this.AgentId],
		"fetch_child":   fetchChild,
	}
	body, err := HttpGet(ReqUrl, ReqQueryHead)
	if err != nil {
		Log("Http Get Error:%v\n", err)
	}
	var simplelist CorpContactUserResponseSimplelist
	err = json.Unmarshal(body.([]byte), &simplelist)
	Log("jsonObj:%v err:%v\n", simplelist, err)
	this.Data["json"] = simplelist
	this.ServeJSON()
	return
}

func (this *ContactAppUserSimplelistController) Post() {
	this.Get()
	return
}
