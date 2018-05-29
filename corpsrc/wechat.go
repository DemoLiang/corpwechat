package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"time"
	"sync"
)

func GetkUserInfo(accessToken,code string)(userid string,err error)  {
	client := &http.Client{}

	req, err := http.NewRequest("Get", "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo",nil)
	if err != nil {
		Log("GetkUserInfo err:%v",err)
		return "",err
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
		fmt.Printf("err:%v",err)
	}
	Log(string(body))
	var userInfo WeChatUserInfo
	json.Unmarshal(body,&userInfo)
	return userInfo.UserId,nil
}


func WeChatAuth(seceret string) string {
	client := &http.Client{}

	req, err := http.NewRequest("Get", "https://qyapi.weixin.qq.com/cgi-bin/gettoken",nil)
	if err != nil {
		Log("WeChatAuth err:%v",err)
		return ""
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	q := req.URL.Query()
	q.Add("corpid", APPID)
	q.Add("corpsecret", seceret)
	req.URL.RawQuery = q.Encode()
	Log(req.URL.String())

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err:%v",err)
	}
	Log("resp.Body:%s\n",body)
	var jsonObj WeChatAuthOutput
	json.Unmarshal(body,&jsonObj)
	Log("jsonObj.AccessToken:%v err:%v\n",jsonObj.AccessToken,err)
	return jsonObj.AccessToken
}


func InitWeChatMap() {
	weChatLock = new(sync.RWMutex)
	weChatLock.Lock()
	for key,seceret := range AgentIdMap {
		AccessTokenMap[key] = WeChatAuth(seceret)
	}
	weChatLock.Unlock()
}

func UpdateWeChatMap() {
	for {
		time.Sleep(time.Minute * 100)
		weChatLock.Lock()
		for key,seceret := range AgentIdMap {
			AccessTokenMap[key] = WeChatAuth(seceret)
		}
		weChatLock.Unlock()
	}
}
