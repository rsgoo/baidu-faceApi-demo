package main

import (
	"baidu-faceApi-demo/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//在线图片活体V3接口文档：https://ai.baidu.com/ai-doc/FACE/Zk37c1urr

//人脸比对v3调用地址：
const FACE_V3_FACEVERIFY_API_URL = "https://aip.baidubce.com/rest/2.0/face/v3/faceverify?access_token="

func main() {
	access_token := utils.GenAccessToken() //获取百度 auth_token
	requestUrl := FACE_V3_FACEVERIFY_API_URL + access_token
	//构建请求参数
	requestJsonString, err := genFaceVerifyRequestJson()
	//异常处理比较简陋
	if err != nil {
		panic(err)
	}

	//构建HTTP请求
	request, err := http.NewRequest(http.MethodPost, requestUrl, strings.NewReader(requestJsonString))
	//异常处理比较简陋
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json")

	//发送HTTP请求
	response, err := http.DefaultClient.Do(request)
	//异常处理比较简陋
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	//读取 response body 内容
	rspBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	fmt.Print(string(rspBytes))

}

func genFaceVerifyRequestJson() (string, error) {
	//获取图片文件base64值
	imageBase64_1, err := utils.GetFileBase64("../resources/images/lenin_1.jpeg")
	imageBase64_2, err := utils.GetFileBase64("../resources/images/lenin_2.jpeg")
	imageBase64_3, err := utils.GetFileBase64("../resources/images/example.jpg")

	//异常处理比较简陋
	if err != nil {
		panic(err)
	}

	//构建请求参数
	arrRequestParams := make([]map[string]interface{}, 3)

	requestParams_1 := make(map[string]interface{})
	requestParams_1["image"] = imageBase64_1
	requestParams_1["image_type"] = "BASE64"
	requestParams_1["face_field"] = "age,beauty,expression,face_shape,gender,glasses,landmark,quality,face_type,spoofing"
	requestParams_1["option"] = "COMMON"

	requestParams_2 := make(map[string]interface{})
	requestParams_2["image"] = imageBase64_2
	requestParams_2["image_type"] = "BASE64"
	requestParams_2["face_field"] = "age,beauty,expression,face_shape,gender,glasses,landmark,quality,face_type,spoofing"
	requestParams_2["option"] = "COMMON"

	requestParams_3 := make(map[string]interface{})
	requestParams_3["image"] = imageBase64_3
	requestParams_3["image_type"] = "BASE64"
	requestParams_3["face_field"] = "age,beauty,expression,face_shape,gender,glasses,landmark,quality,face_type,spoofing"
	requestParams_3["option"] = "COMMON"

	arrRequestParams[0] = requestParams_1
	arrRequestParams[1] = requestParams_2
	arrRequestParams[2] = requestParams_3

	requestBytes, err := json.Marshal(arrRequestParams)
	//异常处理比较简陋
	if err != nil {
		panic(err)
	}
	return string(requestBytes), nil
}
