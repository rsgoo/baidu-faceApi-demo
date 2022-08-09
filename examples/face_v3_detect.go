package main

import (
	"baidu-faceApi-demo/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//人脸检测接口接口文档：https://ai.baidu.com/ai-doc/FACE/yk37c1u4t

//人脸检测接口调用地址
const FACE_V3_DETECT_API_URL = "https://aip.baidubce.com/rest/2.0/face/v3/detect?access_token="

func main() {
	access_token := utils.GenAccessToken() //获取百度 auth_token
	requestUrl := FACE_V3_DETECT_API_URL + access_token

	//获取图片文件base64值
	imageBase64, err := utils.GetFileBase64("../resources/images/example.jpg")

	//异常处理比较简陋
	if err != nil {
		panic(err)
	}

	//构建请求参数
	requestParams := make(map[string]string)
	requestParams["image"] = imageBase64
	requestParams["image_type"] = "BASE64"
	requestParams["face_field"] = "age,beauty,expression,face_shape,gender,glasses,landmark,landmark150,quality,eye_status,emotion,face_type,mask,spoofing"

	requestBytes, err := json.Marshal(requestParams)

	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest(http.MethodPost, requestUrl, bytes.NewBuffer(requestBytes))

	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	response, err := client.Do(request)

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
