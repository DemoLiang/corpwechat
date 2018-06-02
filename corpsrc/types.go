package main

import "encoding/xml"

//corp app 的map结构，用户标记每个应用
type CorpWeChatApp struct {
	EncodingAESKey string `json:"encoding_aes_key"`
	AesKey         string `json:"aes_key"`
	Token          string `json:"token"`
	CoprId         string `json:"copr_id"`
	Seceret        string `json:"seceret"`
}
//获取accesstoken的输出结构
type WeChatAuthOutput struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}
//微信的用户信息
type WeChatUserInfo struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	UserId  string `json:"userid"`
}
//加密的请求内容
type EncryptRequestBody struct {
	XMLName    xml.Name `xml:"xml"`
	ToUserName string   `xml:"ToUserName"`
	Encrypt    string   `xml:"Encrypt"`
	AgentId    string   `xml:"AgentID"`
}

//加密的返回内容
type EncryptResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      string   `xml:"Encrypt"`
	MsgSignature string   `xml:"MsgSignature"`
	TimeStamp    string   `xml:"TimeStamp"`
	Nonce        string   `xml:"Nonce"`
}

//XMl CDATA 数据编码使用
type CDATAText struct {
	Text string `xml:",innerxml"`
}

//基础请求信息
type BasicRequestInfo struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	AgentID      string   `xml:"AgentID"`
}
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
