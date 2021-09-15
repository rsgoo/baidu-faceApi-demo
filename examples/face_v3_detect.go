package main

import (
	"baidu-faceApi-demo/utils"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//百度云接口
const API_URL = "https://aip.baidubce.com/rest/2.0/face/v3/detect?access_token="

//人脸检测：https://ai.baidu.com/ai-doc/FACE/yk37c1u4t
func main() {
	//获取百度 auth_token
	access_token := utils.GenAccessToken()
	requestUrl := API_URL + access_token

	//获取文件base64值
	imageBase64, err := utils.GetFileBase64("../resources/images/example.jpg")

	//异常处理比较简陋
	if err != nil {
		panic(err)
	}

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

	//使用浏览器访问 http://localhost:8088/v3_face_detect 可以得到接口响应结果
	http.HandleFunc("/v3_face_detect", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write(rspBytes)
	})
	
	_ = http.ListenAndServe(":8088", nil)

}
