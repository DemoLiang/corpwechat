package main

import (
	"github.com/astaxie/beego"
	"encoding/base64"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
)

const (
	AgentId = "1000002"
	EncodingAESKey = ""
	Token = ""
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

	if state == "" || code == "" {
		this.Ctx.WriteString("login fail")
		return
	}

	userid, _ := GetkUserInfo(AccessTokenMap[AgentId], code)
	//TODO handler login

	this.Ctx.WriteString("login success userid:" + userid)
	return
}

type WechatServerController struct {
	beego.Controller
}

func (this *WechatServerController) Get() {
	this.Post()
}

func (this *WechatServerController) Post() {
	msgSignature := this.GetString("msg_signature")
	timestamp := this.GetString("timestamp")
	nonce := this.GetString("nonce")
	echostr := this.GetString("echostr")

	Log("msg_signature:%v\n timestamp:%v\n nonce:%v\n echostr:%v\n",msgSignature,timestamp,nonce,echostr)
	if msgSignature == "" || echostr == "" || nonce == ""||timestamp ==""{
		Log("url check fail\n")
		return
	}
	aesMsg,_ := base64.StdEncoding.DecodeString(echostr)
	key,_:=base64.StdEncoding.DecodeString(EncodingAESKey+"=")
	orgMsg, err := AesDecrypt(aesMsg, key)
	if err != nil {
		panic(err)
	}

	// Read length
	buf := bytes.NewBuffer(orgMsg[16:20])
	var length int32
	binary.Read(buf, binary.BigEndian, &length)

	msg := orgMsg[20:20+length]
	this.Ctx.WriteString(string(msg))
	return
}


func EncodeDecodeCheck(){

	msgSignature := ""
	timestamp := ""
	nonce := ""
	echostr := ""

	Log("msg_signature:%v\n timestamp:%v\n nonce:%v\n echostr:%v\n",
		msgSignature,timestamp,nonce,echostr)
	if msgSignature == "" || echostr == "" || nonce == ""||timestamp ==""{
		Log("url check fail\n")
		return
	}
	aesMsg,_ := base64.StdEncoding.DecodeString(echostr)
	key,_:=base64.StdEncoding.DecodeString(EncodingAESKey+"=")
	Log("aesMsg:%X\n key:%X\n len(EncodingAESKey):%v\n",aesMsg,key,len(key))
	origData, err := AesDecrypt(aesMsg, key)
	if err != nil {
		panic(err)
	}
	Log("origData:%s\n",origData)
	Log("origData:%v",origData[15:15+4])
	// Read length
	buf := bytes.NewBuffer(origData[16:20])
	var length int32
	binary.Read(buf, binary.BigEndian, &length)
	//msgLen,_ := strconv.Atoi(string(origData[15:15+4]))
	Log("msgLen:%v\n",length)
	msg := origData[20:20+length]
	Log("msg:%s\n",msg)
	return
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext) % blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}