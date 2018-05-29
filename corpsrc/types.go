package main

type WeChatAuthOutput struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

type WeChatUserInfo struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	UserId  string `json:"userid"`
}
