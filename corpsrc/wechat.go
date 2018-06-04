package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

//获取用户userinfo信息
func GetkUserInfo(accessToken, code string) (userid string, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("Get", "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo", nil)
	if err != nil {
		Log("GetkUserInfo err:%v", err)
		return "", err
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	q.Add("code", code)
	req.URL.RawQuery = q.Encode()

	Log(req.URL.String())

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	Log(string(body))
	var userInfo WeChatUserInfo
	json.Unmarshal(body, &userInfo)
	return userInfo.UserId, nil
}

//请求accesstoken
func WeChatAuth(seceret string) string {
	client := &http.Client{}

	req, err := http.NewRequest("Get", "https://qyapi.weixin.qq.com/cgi-bin/gettoken", nil)
	if err != nil {
		Log("WeChatAuth err:%v", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	q := req.URL.Query()
	q.Add("corpid", CorpID)
	q.Add("corpsecret", seceret)
	req.URL.RawQuery = q.Encode()
	Log(req.URL.String())

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	Log("resp.Body:%s\n", body)
	var jsonObj WeChatAuthOutput
	json.Unmarshal(body, &jsonObj)
	Log("jsonObj.AccessToken:%v err:%v\n", jsonObj.AccessToken, err)
	return jsonObj.AccessToken
}

//生成签名
func MakeMsgSignature(token, timestamp, nonce, msgencrypt string) string {
	sl := []string{token, timestamp, nonce, msgencrypt}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%02x", s.Sum(nil))
}

//验证数据是否正确
func ValidateMsg(token, timestamp, nonce, msgEncrypt, msgSignatureIn string) bool {
	msgSignatureGen := MakeMsgSignature(token, timestamp, nonce, msgEncrypt)
	if msgSignatureGen != msgSignatureIn {
		Log("msgSignatureGen:%v msgSignatureIn:%v\n", msgSignatureGen, msgSignatureIn)
		return false
	}
	return true
}

//XML解码基础请求到加密结构
func ParseEncryptRequestBody(body []byte) *EncryptRequestBody {
	requestBody := &EncryptRequestBody{}
	xml.Unmarshal(body, requestBody)
	return requestBody
}

//AES Key 做base64解码
func EncodingAESKey2AESKey(encodingAESKey string) []byte {
	data, _ := base64.StdEncoding.DecodeString(encodingAESKey + "=")
	return data
}

//编码明文内容得到msg_encrypt
func MakeEncryptXmlData(body string, AesKey string) (string, error) {
	// Encrypt part2: Length bytes
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, int32(len(body)))
	if err != nil {
		Log("Binary write err:\n", err)
	}
	bodyLength := buf.Bytes()

	// Encrypt part1: Random bytes
	randomBytes := []byte(GenRandomString(16))

	// Encrypt Part, with part4 - appID
	plainData := bytes.Join([][]byte{randomBytes, bodyLength, []byte(body), []byte(CorpID)}, nil)
	cipherData, err := AesEncrypt(plainData, []byte(AesKey))
	if err != nil {
		return "", errors.New("aesEncrypt error")
	}

	return base64.StdEncoding.EncodeToString(cipherData), nil
}

//解码解密内容为明文内容，并拿取实际的content
func MakeDecryptXMLData(encryptBody string, AesKey string) (string, error) {
	aesMsg, _ := base64.StdEncoding.DecodeString(encryptBody)
	orgMsg, err := AesDecrypt(aesMsg, []byte(AesKey))
	if err != nil {
		panic(err)
	}

	// Read length
	buf := bytes.NewBuffer(orgMsg[16:20])
	var length int32
	binary.Read(buf, binary.BigEndian, &length)

	msg := orgMsg[20 : 20+length]
	return string(msg), nil
}

//生成XML CDATA 字符
func Value2CDATA(v string) CDATAText {
	return CDATAText{"<![CDATA[" + v + "]]>"}
}

//生成随机字符串
func GenRandomString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//XML编码响应包
func MakeEncryptResponseBody(token, content, nonce, timestamp string, AesKey string) ([]byte, error) {
	encryptBody := &EncryptResponseBody{}

	encryptXmlData, _ := MakeEncryptXmlData(content, AesKey)
	encryptBody.Encrypt = encryptXmlData
	encryptBody.MsgSignature = MakeMsgSignature(token, timestamp, nonce, encryptXmlData)
	encryptBody.TimeStamp = timestamp
	encryptBody.Nonce = nonce

	return xml.MarshalIndent(encryptBody, " ", "  ")
}

//XML解析到文本内容结构
func ParseEncryptTextRequestBody(plainText []byte) (*TextRequestMessage, error) {
	textRequestBody := &TextRequestMessage{}
	xml.Unmarshal(plainText, textRequestBody)
	return textRequestBody, nil
}

//XML解析到基础结构
func ParseEncryptBasicRequestInfo(plainText []byte) (*BasicRequestInfo, error) {
	basicRequestInfo := &BasicRequestInfo{}
	xml.Unmarshal(plainText, &basicRequestInfo)
	return basicRequestInfo, nil
}

//XML 解析到通讯录便跟世界爱你
func ParseEncryptEventChangeContact(plainText []byte) (*ChangeContactRequestEventMessage, error) {
	changeContact := &ChangeContactRequestEventMessage{}
	xml.Unmarshal(plainText, &changeContact)
	return changeContact, nil
}

//读取解密后的内容
func ParseEncryptRequestBodyContent(plainText []byte) ([]byte, error) {
	// Read length
	buf := bytes.NewBuffer(plainText[16:20])
	var length int32
	binary.Read(buf, binary.BigEndian, &length)
	// appID validation
	appIDstart := 20 + length
	id := plainText[appIDstart : int(appIDstart)+len(CorpID)]
	if !ValidateAppId(id) {
		Log("Wechat Message Service: appid is invalid!")
		return nil, errors.New("Appid is invalid")
	}
	Log("Wechat Message Service: appid validation is ok!\n")

	return plainText[20 : 20+length], nil
}

//验证corpid
func ValidateAppId(id []byte) bool {
	if string(id) == CorpID {
		return true
	}
	return false
}

func HttpGet(ReqUrl string, ReqQueryHead map[string]string) (interface{}, error) {
	client := &http.Client{}

	req, err := http.NewRequest("Get", ReqUrl, nil)
	if err != nil {
		Log("WeChatAuth err:%v", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	q := req.URL.Query()
	for key, data := range ReqQueryHead {
		q.Add(key, data)
	}
	req.URL.RawQuery = q.Encode()
	Log(req.URL.String())

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	Log("resp.Body:%s\n", body)
	return body, nil
}

//初始化Token
func InitWeChatMap() {
	weChatLock = new(sync.RWMutex)
	weChatLock.Lock()
	for key, corpApp := range CorpAppMap {
		AccessTokenMap[key] = WeChatAuth(corpApp.Seceret)
	}
	weChatLock.Unlock()
}

//定期更新token
func UpdateWeChatMap() {
	for {
		time.Sleep(time.Minute * 100)
		weChatLock.Lock()
		for key, corpApp := range CorpAppMap {
			AccessTokenMap[key] = WeChatAuth(corpApp.Seceret)
		}
		weChatLock.Unlock()
	}
}
