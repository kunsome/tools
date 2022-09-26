package main

import (
	"fmt"
	"github.com/kunsome/tools/console"
)

type MyParam struct {
	SomePath string `json:"路径参数" inputType:"dir"`
	SomeInt  int    `json:"数字参数"`
	SomeStr  string `json:"字符串参数" `
}

func main() {
	var p MyParam
	console.CircleHelper(func() {
		console.LoadParam(&p)
		fmt.Println("根据参数执行逻辑")
	})

}
