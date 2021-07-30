package utils

import (
	"fmt"
	"io/ioutil"
)

const AUTH_URL = "https://aip.baidubce.com/oauth/2.0/token"

type AllConf struct{}

type AuthConf struct {
	Grant_type    string `json:"grant_type"`
	Client_id     string `json:"client_id"`
	Client_secret string `json:"client_secret"`
}

//获取 AccessToken
func GetAccessToken() {
	bytes, err := ioutil.ReadFile("../conf/auth_account.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bytes))
}
