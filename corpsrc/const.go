package main

const (
	CorpID = "ww73b5f22913f98256"
)

const (
	REQ__MSG_TYPE__TEXT                 = "text"
	REQ__MSG_TYPE__IMAGE                = "image"
	REQ__MSG_TYPE__VOICE                = "voice"
	REQ__MSG_TYPE__VIDEO                = "video"
	REQ__MSG_TYPE__LOCATION             = "location"
	REQ__MSG_TYPE__LINK                 = "link"
	REQ__EVENT_TYPE__SUBSCRIBE          = "subscribe"
	REQ__EVENT_TYPE__ENTER_AGENT        = "enter_agent"
	REQ__EVENT_TYPE__LOCATION           = "LOCATION"
	REQ__EVENT_TYPE__BATCH_JOB_RESULT   = "batch_job_result"
	REQ__EVENT_TYPE__CHANGE_CONTACT     = "change_contact"
	REQ__EVENT_TYPE__UPDATE_USER        = "update_user"
	REQ__EVENT_TYPE__DELETE_USER        = "delete_user"
	REQ__EVENT_TYPE__CREATE_PARTY       = "create_party"
	REQ__EVENT_TYPE__UPDATE_PARTY       = "update_party"
	REQ__EVENT_TYPE__DELETE_PARTY       = "delete_party"
	REQ__EVENT_TYPE__UPDATE_TAG         = "update_tag"
	REQ__EVENT_TYPE__CLICK              = "click"
	REQ__EVENT_TYPE__VIEW               = "view"
	REQ__EVENT_TYPE__SCANCODE_PUSH      = "scancode_push"
	REQ__EVENT_TYPE__SCANCODE_WAITMSG   = "scancode_waitmsg"
	REQ__EVENT_TYPE__PIC_SYSPHOTO       = "pic_sysphoto"
	REQ__EVENT_TYPE__PIC_PHOTO_OR_ALBUM = "pic_photo_or_album"
	REQ__EVENT_TYPE__PIC_WEIXIN         = "pic_weixin"
	REQ__EVENT_TYPE__LOCATION_SELECT    = "location_select"
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
