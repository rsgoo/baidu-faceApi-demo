package utils

import (
	"encoding/base64"
	"io/ioutil"
)

func GetFileBase64(filePath string) (string, error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	fileBase64 := base64.StdEncoding.EncodeToString(fileBytes)
	return fileBase64, nil
}
