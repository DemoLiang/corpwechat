package main

const (
	CorpID = "ww73b5f22913f98256"
)

const (
	REQ__MSG_TYPE__TEXT = "text"
)

var CorpAppMap = map[string]CorpWeChatApp{
	//"XXXX": CorpWeChatApp{
	//	EncodingAESKey: "XXXXX",
	//	Token:          "XXXX",
	//	CoprId:         "1000002",
	//	Seceret:        "XXXXX",
	//},
}

func init() {
	for key, corpApp := range CorpAppMap {
		corpApp.AesKey = string(EncodingAESKey2AESKey(corpApp.EncodingAESKey))
		CorpAppMap[key] = corpApp
	}
}
