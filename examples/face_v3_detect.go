package main

import (
	"baidu-faceApi-demo/utils"
	"fmt"
)

func main() {
	access_token := utils.GenAccessToken()
	fmt.Println(access_token)
}
