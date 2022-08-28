package main

import (
	"baidu-faceApi-demo/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//人脸对比V3接口文档：https://ai.baidu.com/ai-doc/FACE/Lk37c1tpf

//人脸比对v3调用地址：
const FACE_V3_MATCH_API_URL = "https://aip.baidubce.com/rest/2.0/face/v3/match?access_token="

//秋
func main() {
	access_token := utils.GenAccessToken()
	requestUrl := FACE_V3_MATCH_API_URL + access_token

	requestJsonString, err := genMatchRequestJson()

	if err != nil {
		panic(err)
	}

	//进行 http 请求
	request, err := http.NewRequest(http.MethodPost, requestUrl, strings.NewReader(requestJsonString))

	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	//读取 response 内容
	rspBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Print(string(rspBytes))
}

func genMatchRequestJson() (string, error) {
	//获取图片文件base64值
	imageBase64_1, err := utils.GetFileBase64("../resources/images/lenin_1.jpeg")
	imageBase64_2, err := utils.GetFileBase64("../resources/images/lenin_2.jpeg")

	//异常处理比较简陋
	if err != nil {
		panic(err)
	}

	//构建请求参数
	arrRequestParams := make([]map[string]interface{}, 2)

	requestParams_1 := make(map[string]interface{})
	requestParams_1["image"] = imageBase64_1
	requestParams_1["image_type"] = "BASE64"
	requestParams_1["face_type"] = "LIVE"
	requestParams_1["quality_control"] = "LOW"
	requestParams_1["liveness_control"] = "NONE"
	requestParams_1["spoofing_control"] = "LOW"
	requestParams_1["face_sort_type"] = 1

	requestParams_2 := make(map[string]interface{})
	requestParams_2["image"] = imageBase64_2
	requestParams_2["image_type"] = "BASE64"
	requestParams_2["face_type"] = "LIVE"
	requestParams_2["quality_control"] = "LOW"
	requestParams_2["liveness_control"] = "LOW"
	requestParams_2["spoofing_control"] = "LOW"
	requestParams_2["face_sort_type"] = 1

	arrRequestParams[0] = requestParams_1
	arrRequestParams[1] = requestParams_2

	requestBytes, err := json.Marshal(arrRequestParams)

	//异常处理比较简陋
	if err != nil {
		panic(err)
	}
	return string(requestBytes), nil
}
