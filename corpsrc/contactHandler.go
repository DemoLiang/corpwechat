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
	"strings"
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
	return
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
