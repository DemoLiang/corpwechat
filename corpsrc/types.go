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
	Event        string   `xml:"Event"`
	MsgId        int64    `xml:"MsgId"`
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

//图片消息请求body content 内容
type ImageRequestMessage struct {
	BasicRequestInfo
	PicUrl   string `xml:"PicUrl"`
	MedialId string `xml:"MedialId"`
}

//图片消息回复消息格式
type ImageResponseMessage struct {
	BasicResponseInfo
	Image struct {
		MediaId string `xml:"MediaId"`
	} `xml:"Image"`
}

//图片消息响应媒体结构
//type ImageResponseImage struct {
//	MediaId string `xml:"MediaId"`
//}

//语音消息请求body content 内容
type VoiceRequestMessage struct {
	BasicRequestInfo
	MediaId string `xml:"MediaId"`
	Format  string `xml:"Format"`
}

//语音消息响应消息格式
type VoiceResponseMessage struct {
	BasicResponseInfo
	Voice struct {
		MediaId string `xml:"MediaId"`
	} `xml:"Voice"`
}

//视频消息请求body
type VideoRequestMessage struct {
	BasicRequestInfo
	MediaId      string `xml:"MediaId"`
	ThumbMediaId string `xml:"ThumbMediaId"`
}

//视频消息响应格式
type VideoResponseMessage struct {
	BasicResponseInfo
	Video struct {
		MediaId     string `xml:"MediaId"`
		Title       string `xml:"Title"`
		Description string `xml:"Description"`
	} `xml:"Video"`
}

//图文消息响应格式
type NewsResponseMessage struct {
	BasicResponseInfo
	ArticleCount int `xml:"ArticleCount"`
	Articles     struct {
		Item []NewsItem `xml:"Item"`
	} `xml:"Articles"`
}

//图文列表结构
type NewsItem struct {
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	PicUrl      string `xml:"PicUrl"`
	Url         string `xml:"Url"`
}

//位置消息请求body
type LocationRequestMessage struct {
	BasicRequestInfo
	LocationX float64 `xml:"Location_X"`
	LocationY float64 `xml:"Location_Y"`
	Scale     int     `xml:"Scale"`
	Label     string  `xml:"Label"`
}

//链接消息请求body
type LinkRequestMessage struct {
	BasicRequestInfo
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	PicUrl      string `xml:"PicUrl"`
}

//关注消息事件请求body
type SubscribeRequestEventMessage struct {
	BasicRequestInfo
}

//进入应用事件消息
type EnterAgentRequestEventMessage struct {
	BasicRequestInfo
	EventKey string `xml:"EventKey"`
}

//上报地址位置事件消息
type UploadLocationRequestEventMessage struct {
	BasicRequestInfo
	Latitude  float64 `xml:"Latitude"`
	Longitude float64 `xml:"Longitude"`
	Precision float64 `xml:"Precision"`
}

//异步任务完成事件推送
type BatchJobResultRequestEventMessage struct {
	BasicRequestInfo
	BatchJob BatchJobEventBody `xml:"BatchJob"`
}

//异步事件结构
type BatchJobEventBody struct {
	JobId   string `xml:"JobId"`
	JobType string `xml:"JobType"`
	ErrCode int    `xml:"ErrCode"`
	ErrMsg  string `xml:"ErrMsg"`
}

//通讯录变更请求事件消息
type ChangeContactRequestEventMessage struct {
	BasicRequestInfo
	ChangeType  string         `xml:"ChangeType"`
	UserID      string         `xml:"UserID"`
	Name        string         `xml:"Name"`
	Department  string         `xml:"Department"`
	Position    string         `xml:"Position"`
	Mobile      string         `xml:"Mobile"`
	Gender      int            `xml:"Gender"`
	Email       string         `xml:"Email"`
	Status      int            `xml:"Status"`
	Avatar      string         `xml:"Avatar"`
	EnglishName string         `xml:"EnglishName"`
	IsLeader    int            `xml:"IsLeader"`
	Telephone   string         `xml:"Telephone"`
	ExtAttr     ContactExtAttr `xml:"ExtAttr"`
}

//更新成员请求事件
type UpdateContactUseRequestEventMessage struct {
	BasicRequestInfo
	ChangeType  string         `xml:"ChangeType"`
	UserID      string         `xml:"UserID"`
	NewUserID   string         `xml:"NewUserID"`
	Department  string         `xml:"Department"`
	Position    string         `xml:"Position"`
	Mobile      string         `xml:"Mobile"`
	Gender      int            `xml:"Gender"`
	Email       string         `xml:"Email"`
	Status      int            `xml:"Status"`
	Avatar      string         `xml:"Avatar"`
	EnglishName string         `xml:"EnglishName"`
	IsLeader    int            `xml:"IsLeader"`
	Telephone   string         `xml:"Telephone"`
	ExtAttr     ContactExtAttr `xml:"ExtAttr"`
}

//删除成员请求事件
type DeleteContactUserRequestEventMessage struct {
	BasicRequestInfo
	ChangeType string `xml:"ChangeType"`
	UserID     string `xml:"UserID"`
}

//通讯录扩展属性
type ContactExtAttr struct {
	Item []ContactExtAttrItem `xml:"Item"`
}

//通讯录扩展数据项
type ContactExtAttrItem struct {
	Name  string `xml:"Name"`
	Value string `xml:"Value"`
}

//新增部门请求事件
type CreatePartyRequestEventMessage struct {
	BasicRequestInfo
	ChangeType string `xml:"ChangeType"`
	Id         int    `xml:"Id"`
	Name       string `xml:"Name"`
	ParentId   string `xml:"ParentId"`
	Order      int    `xml:"Order"`
}

//更新部门请求事件
type UpdatePartyRequestEventMessage struct {
	BasicRequestInfo
	ChangeType string `xml:"ChangeType"`
	Id         int    `xml:"Id"`
	Name       string `xml:"Name"`
	ParentId   string `xml:"ParentId"`
}

//删除部门请求事件
type DeletePartyRequestEventMessage struct {
	BasicRequestInfo
	ChangeType string `xml:"ChangeType"`
	Id         int    `xml:"Id"`
}

//标签成员变更请求事件
type UpdateTagRequestEventMessage struct {
	BasicRequestInfo
	ChangeType    string `xml:"ChangeType"`
	TagId         int    `xml:"TagId"`
	AddUserItems  string `xml:"AddUserItems"`
	DelUserItems  string `xml:"DelUserItems"`
	AddPartyItems string `xml:"AddPartyItems"`
	DelPartyItems string `xml:"DelPartyItems"`
}

//菜单点击请求事件
type MenuClickRequestEventMessage struct {
	BasicRequestInfo
	EventKey string `xml:"EventKey"`
}

//菜单点击跳转到链接事件
type MenuClickRedirectRequesstEventMessage struct {
	BasicRequestInfo
	EventKey string `xml:"EventKey"`
}

//扫码推事件
type QrcodeRequestEventMessage struct {
	BasicRequestInfo
	EventKey     string       `xml:"EventKey"`
	ScanCodeInfo ScanCodeInfo `xml:"ScanCodeInfo"`
}

//扫码推事件且弹出消息接收中事件推送
type ScancodeWaitmsgRequestEventMessage struct {
	BasicRequestInfo
	EventKey     string
	ScanCodeInfo ScanCodeInfo `xml:"ScanCodeInfo"`
}

//扫描信息
type ScanCodeInfo struct {
	ScanType   string `xml:"ScanType"`
	ScanResult string `xml:"ScanResult"`
}

//弹出系统拍照发图事件推送
type PicSysphotoRequestEventMessage struct {
	BasicRequestInfo
	EventKey     string                  `xml:"EventKey"`
	SendPicsInfo PicSysphotoSendPicsInfo `xml:"SendPicsInfo"`
}

//弹出拍照或者相册发图的事件推送
type PicPhotoOrAlbumRequestEventMessage struct {
	BasicRequestInfo
	EventKey     string                  `xml:"EventKey"`
	SendPicsInfo PicSysphotoSendPicsInfo `xml:"SendPicsInfo"`
}

//弹出微信相册发图器的事件推送
type PicWeixinRequestEventMessage struct {
	BasicRequestInfo
	EventKey     string                  `xml:"EventKey"`
	SendPicsInfo PicSysphotoSendPicsInfo `xml:"SendPicsInfo"`
}

//发送图片的信息
type PicSysphotoSendPicsInfo struct {
	Count   int                  `xml:"Count"`
	PicList []PicSysphotoPicList `xml:"PicList"`
}

//图片列表
type PicSysphotoPicList struct {
	Item PicSysphotoPicItem `xml:"Item"`
}

//图片MD5信息
type PicSysphotoPicItem struct {
	PicMd5Sum string `xml:"PicMd5Sum"`
}

//弹出地理位置选择器的事件推送
type LocationSelectRequestEventMessage struct {
	BasicRequestInfo
	EventKey         string                         `xml:"EventKey"`
	SendLocationInfo LocationSelectSendLocationInfo `xml:"SendLocationInfo"`
}

//发送的位置信息
type LocationSelectSendLocationInfo struct {
	LocationX float64 `xml:"Location_X"`
	LocationY float64 `xml:"Location_Y"`
	Scale     int     `xml:"Scale"`
	Label     string  `xml:"Label"`
	Poiname   string  `xml:"Poiname"`
}

//群聊消息基础响应结构
type AppchatBasicResponseMessage struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

//创建群聊会话
type AppchatCreateRequestMessage struct {
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	Userlist []string `json:"userlist"`
	Chatid   string   `json:"chatid"`
}

//创建群聊响应消息
type AppchatCreateResponseMessage struct {
	AppchatBasicResponseMessage
	Chatid string `json:"chatid"`
}

//修改群聊会话
type AppchatUpdateRequestMessage struct {
	Chatid      string   `json:"chatid"`
	Name        string   `json:"name"`
	Owner       string   `json:"owner"`
	AddUserList []string `json:"add_user_list"`
	DelUserList []string `json:"del_user_list"`
}

//获取群聊会话
type AppchatGetResponsetMessage struct {
	AppchatBasicResponseMessage
	ChatInfo AppchatGetChatInfo `json:"chat_info"`
}

//群聊会话信息
type AppchatGetChatInfo struct {
	Chatid   string   `json:"chatid"`
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	Userlist []string `json:"userlist"`
}

//群聊消息推送基础结构
type AppchatSendBasic struct {
	Chatid  string `json:"chatid"`
	Msgtype string `json:"msgtype"`
	Safe    int    `json:"safe"`
}

//群聊消息 文本消息推送
type AppchatSendText struct {
	AppchatSendBasic
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

//群聊消息 图片消息推送
type AppchatSendImage struct {
	AppchatSendBasic
	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

//语音消息
type AppchatSendVoice struct {
	AppchatSendBasic
	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
}

//视频消息
type AppchatSendVideo struct {
	AppchatSendBasic
	Video struct {
		MediaId     string `json:"media_id"`
		Description string `json:"description"`
	} `json:"video"`
}

//文件消息
type AppchatSendFile struct {
	AppchatSendBasic
	File struct {
		MediaId string `json:"media_id"`
	} `json:"file"`
}

//文本卡消息
type AppchatSendTextCard struct {
	AppchatSendBasic
	Textcard struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Url         string `json:"url"`
		Btntxt      string `json:"btntxt"`
	} `json:"textcard"`
}

//图文消息推送News
type AppchatSendNews struct {
	AppchatSendBasic
	News struct {
		Articles []AppchatSendNewsArticles `json:"articles"`
	} `json:"news"`
}

//news文章结构
type AppchatSendNewsArticles struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Picurl      string `json:"picurl"`
	Btntxt      string `json:"btntxt"`
}

//图文消息推送mpnews
type AppchatSendMpnews struct {
	AppchatSendBasic
	Mpnews struct {
		Articles []AppchatSendMpnewsArticles `json:"articles"`
	} `json:"mpnews"`
}

//mpnews文章结构
type AppchatSendMpnewsArticles struct {
	Title            string `json:"title"`
	ThumbMediaId     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	ContentSourceUrl string `json:"content_source_url"`
	Content          string `json:"content"`
	Digest           string `json:"digest"`
}

//发送应用消息
type MessageSendResponseBasic struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Invaliduser  string `json:"invaliduser"`
	Invalidparty string `json:"invalidparty"`
	Invalidtag   string `json:"invalidtag"`
}

//发送应用消息基础结构
type MessageSendBasic struct {
	Touser  string `json:"touser"`
	Toparty string `json:"toparty"`
	Totag   string `json:"totag"`
	Msgtype string `json:"msgtype"`
	Agentid int    `json:"agentid"`
}

//发送应用文本消息
type MessageSendText struct {
	MessageSendBasic
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

//发送应用图片消息
type MessageSendImage struct {
	MessageSendBasic
	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

//发送应用图片消息
type MessageSendVoice struct {
	MessageSendBasic
	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
}

//发送应用视频消息
type MessageSendVideo struct {
	MessageSendBasic
	Video struct {
		MediaId     string `json:"media_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"video"`
}

//发送应用文件消息
type MessageSendFile struct {
	MessageSendBasic
	File struct {
		MediaId string `json:"media_id"`
	} `json:"file"`
}

//发送应用文本卡消息
type MessageSendTextCard struct {
	MessageSendBasic
	Textcard struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Url         string `json:"url"`
		Btntxt      string `json:"btntxt"`
	} `json:"textcard"`
}

//发送应用图文消息News
type MessageSendNews struct {
	MessageSendBasic
	News struct {
		Articles []AppchatSendNewsArticles `json:"articles"`
	} `json:"news"`
}

//发送应用图文消息mpnews
type MessageSendMpnews struct {
	MessageSendBasic
	Mpnews struct {
		Articles []AppchatSendMpnewsArticles `json:"articles"`
	} `json:"mpnews"`
}
