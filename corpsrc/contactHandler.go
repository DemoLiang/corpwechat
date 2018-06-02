package main

import (
	"github.com/astaxie/beego"
	"sort"
	"strings"
	"crypto/sha1"
	"fmt"
	"encoding/base64"
	"bytes"
	"encoding/binary"
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
		orgMsg, err := AesDecrypt(aesMsg,[]byte(CorpAppMap[this.AgentId].AesKey))
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

