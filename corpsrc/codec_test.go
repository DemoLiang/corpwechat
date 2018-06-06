package main

import "testing"

//AES加解密测试
func TestAesEncrypt(t *testing.T) {
	msg := `<xml><ToUserName><![CDATA[LiangGuiMing]]></ToUserName><FromUserName><![CDATA[ww73b5f22913f98256]]></FromUserName><CreateTime>1527682730</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[b]]></Content></xml>`
	aesKey := string(EncodingAESKey2AESKey("XXXX"))
	msgEncrypt, _ := AesEncrypt([]byte(msg), []byte(aesKey))

	msgDecrypt, _ := AesDecrypt(msgEncrypt, []byte(aesKey))
	if string(msgDecrypt) == msg {
		t.Log("success")
	} else {
		t.Log("fail")
	}
}
