package console

// console 包， 命令行帮助函数。
// 创建人： kunsome
// 创建时间e： 20220926

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// GetInputDir ， 从命令行输入中获取一个，正确的文件夹地址
// 参数：
//   - prompt  输入提示信息
//
// 返回值：
//   - path 文件夹地址
func GetInputDir(prompt string) (path string) {
	for path == "" {
		inputStr := ""
		fmt.Println(prompt)
		_, err := fmt.Scanln(&inputStr)
		if err != nil {
			log.Println("获取输入参数错误：", err)
		}
		inputStr = strings.TrimRight(inputStr, string(os.PathSeparator))
		if _, err := os.Stat(inputStr); os.IsNotExist(err) {
			fmt.Println("输入的路径不存在, 请重新输入！！！！！！")
			continue
		}
		path = inputStr
	}
	return
}

// GetInput 命令行输入中获取一个字符串输入
// 参数：
//   - prompt  输入提示信息
//
// 返回值：
//   - inputStr: 字符串输入
func GetInput(prompt string) (inputStr string) {
	fmt.Println(prompt)
	var _, err = fmt.Scanln(&inputStr)
	if err != nil {
		log.Println("获取输入参数错误：", err)
	}
	return inputStr
}

// GetInputInt 命令行输入中获取一个int输入
// 参数：
//   - prompt ： 输入提示信息
//
// 返回值：
//   - inputStr: 字符串输入
func GetInputInt(prompt string) int {
	for {
		ret, err := strconv.Atoi(GetInput(prompt))
		if err != nil {
			log.Println("整形数字输入", err)
			continue
		}
		return ret
	}
}
