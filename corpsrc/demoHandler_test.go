package main

import (
	"testing"
	"encoding/base64"
	"encoding/xml"
	"bytes"
	"encoding/binary"
)

//加解密测试
func TestMakeEncryptXmlData(t *testing.T) {
	body:=`event`
	aesKey:="abc"

	encryptMsg,_ := MakeEncryptXmlData(body,aesKey)
	decryptMsg ,_:=MakeDecryptXMLData(encryptMsg,aesKey)
	if body != decryptMsg{
		t.Log("test fail!")
	}else {
		t.Log(decryptMsg)
		t.Log("test success!")
	}
}

//随机字符串生成测试
func TestGenRandomString(t *testing.T) {
	randstr :=GenRandomString(16)
	t.Log(len(randstr))
}

//文本内容加解密请求返回测试
func TestTextMessageHandler(t *testing.T) {
	xmlMsg := `<xml><ToUserName>ww73b5f22913f98256</ToUserName><FromUserName>ww73b5f22913f98256</FromUserName><CreateTime>1527909705</CreateTime><MsgType>text</MsgType><Content>b</Content></xml>`
	xmlResp:=`<xml><Encrypt>Z7MS2hf41LJ8adRJ1rrKaXHMfzgXPlxn/xBAnfjtXSuhrmP4gWqzUTc8Fn/KVTPsS0O2wfhq6xkQUgEvpVXcl11FE09/nuDvgLIuD0V7EUpxW2HOiW9PWETjBHJCMb6t7o5TLtUhxhE0LjGlBmr60CnB/ZqSrek05oVcxdvh+Z4v3/+Xe4/1PDmZEDrNMRFSHgACj/8TSOi11zrquOGD/KA+mpKWgJgbC2nez1KodTL50OcPx+CQ70KscoE03Go2tlTorrhPhnce44UCt0GY59V3xp6JQi7w9vw0Ta0NfkY=</Encrypt><MsgSignature>db46dc7890d4b3008821eb2f0f40890a11860984</MsgSignature><TimeStamp>1527909705</TimeStamp><Nonce>8DqdwgcI</Nonce></xml>`

	aesKey:=string(EncodingAESKey2AESKey("XXXX"))	//需要填入响应额encodingAESKey

	var respBody EncryptResponseBody
 	xml.Unmarshal([]byte(xmlResp),&respBody)
	t.Log(respBody)
	t.Log(respBody.Encrypt)
	aesMsg, _ := base64.StdEncoding.DecodeString(respBody.Encrypt)
	t.Log(string(aesMsg))
	orgMsg, _ := AesDecrypt(aesMsg, []byte(aesKey))
	t.Log(string(orgMsg))
	buf := bytes.NewBuffer(orgMsg[16:20])
	var length int32
	binary.Read(buf, binary.BigEndian, &length)

	msg := orgMsg[20 : 20+length]

	if xmlMsg != string(msg){
		t.Log("test fail!")
	}else {
		t.Log(string(msg))
		t.Log("test success!")
	}
}