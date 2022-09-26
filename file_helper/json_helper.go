package file_helper

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

// ReadJson 读取json文件并反序列化到target
func ReadJson(filePath string, target interface{}) (err error) {
	var byteValue []byte
	byteValue, err = ReadFile(filePath)
	if err != nil {
		return
	}
	//原始文件翻译成json
	err = json.Unmarshal(byteValue, target)
	if err != nil {
		return
	}
	return
}

// WriteJson json写入文件
func WriteJson(filepath string, obj interface{}) error {
	output, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath, output, 0644)
}

// ReadRemoteJson 读取远程json
func ReadRemoteJson(url string, target interface{}) (err error) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, target)
	return
}
