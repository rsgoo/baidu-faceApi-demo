package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//百度云access_token 生成请求地址
const AUTH_URL = "https://aip.baidubce.com/oauth/2.0/token"

//账户配置-all
type AuthAccount struct {
	AkSkConf AkSkConfig `json:"ak_sk_conf"`
}

//账户配置-auth
type AkSkConfig struct {
	Grant_type    string `json:"grant_type"`
	Client_id     string `json:"client_id"`
	Client_secret string `json:"client_secret"`
}

type RspAuthInfo struct {
	Refresh_token  string `json:"refresh_token"`
	Expires_in     int    `json:"expires_in"`
	Access_token   string `json:"access_token"`
	Scope          string `json:"scope"`
	Session_secret string `json:"session_secret"`
}

//获取Json文件中百度云 ak sk 配置信息
func GetAuthAccount() (result AkSkConfig) {
	confFileName := "../conf/auth_account.json"
	bytes, err := ioutil.ReadFile(confFileName)
	if err != nil {
		fmt.Println("读取配置文件 " + confFileName + "失败,请确认文件是否存在")
		fmt.Println(err.Error())
		return
	}

	akSkInfo := AuthAccount{}
	err = json.Unmarshal(bytes, &akSkInfo)
	if err != nil {
		fmt.Println("配置文件解析失败，请确文件内容是否是有效的json")
		fmt.Println(err.Error())
		return
	}
	result = akSkInfo.AkSkConf
	return result
}

//获取access_token
//todo 根据 access_token 的有效期设置缓存
func GenAccessToken() string {
	authAccount := GetAuthAccount()

	authAccountMap := make(map[string]string)
	authAccountMap["grant_type"] = authAccount.Grant_type
	authAccountMap["client_id"] = authAccount.Client_id
	authAccountMap["client_secret"] = authAccount.Client_secret

	reqUrl := AUTH_URL + "?"
	for key, val := range authAccountMap {
		reqUrl = reqUrl + key + "=" + val + "&"
	}
	rsp, err := http.Post(reqUrl, "", nil)
	if err != nil {
		fmt.Println("请求access_token失败")
		fmt.Println(err.Error())
		return ""
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println("读取请求响应结果失败")
		fmt.Println(err.Error())
		return ""
	}

	response := &RspAuthInfo{}
	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("解析access_token信息失败")
		fmt.Println(err.Error())
		return ""
	}

	if len(response.Access_token) == 0 {
		fmt.Println("请求access_token信息失败,请求检测ak,sk是否正确")
		return ""
	}

	return response.Access_token
}
