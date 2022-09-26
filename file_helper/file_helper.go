package file_helper

import (
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// GetExt 获取文件后缀名
func GetExt(httpUrl string) (ext string, err error) {
	u, err := url.Parse(httpUrl)
	if err != nil {
		log.Println("csv line 图片后缀解析错误", err)
		return
	}
	//赋值后缀名字段
	ext = filepath.Ext(u.Path)
	return
}

// GetFileNameNoExt 获取无后缀的文件名
func GetFileNameNoExt(path string) string {
	name := filepath.Base(path)
	return strings.TrimSuffix(name, filepath.Ext(name))
}

// ReadFile 读取文件内，并返回对应[]byte
func ReadFile(filePath string) (ret []byte, err error) {
	//读取源文件json
	var file *os.File
	file, err = os.Open(filePath)
	defer file.Close()
	if err != nil {
		return
	}
	ret, err = ioutil.ReadAll(file)
	return
}

// IsFileExist 验证文件是否存在
func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

// CreateDir 创建文件夹
func CreateDir(path string) error {
	if !IsFileExist(path) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
