package console

import (
	"github.com/kunsome/tools/print_helper"
	"reflect"
)

func LoadParam(param interface{}) {
	for {
		paramType := reflect.TypeOf(param)
		paramValue := reflect.ValueOf(param)
		for i := 0; i < paramType.Elem().NumField(); i++ {
			fieldNameCn := paramType.Elem().Field(i).Tag.Get("json")
			InputType := paramType.Elem().Field(i).Tag.Get("inputType")
			fileType := paramType.Elem().Field(i).Type.Name()

			if InputType == "dir" {
				paramValue.Elem().Field(i).Set(reflect.ValueOf(GetInputDir("请输入" + fieldNameCn + ":")))
			} else if fileType == "int" {
				paramValue.Elem().Field(i).Set(reflect.ValueOf(GetInputInt("请输入" + fieldNameCn + "：")))
			} else {
				paramValue.Elem().Field(i).Set(reflect.ValueOf(GetInput("请输入" + fieldNameCn + "：")))
			}
		}

		print("\n将以下列参数执行：")
		print_helper.PrintStruct(param)
		quit := GetInput("请确认是否要执行？输入'y'确认：")
		if quit == "y" || quit == "Y" {
			break
		}
	}
}

func BaseInterpreter() {

}

// CircleHelper 循环询问是否退出当前函数,防止点击的执行的命令行程序直接退出，看不到返回结果
func CircleHelper(f func()) {
	for {
		f()
		//处理是否继续执行的相关逻辑
		quit := GetInput("是否继续执行一下一个任务，退出输入x,继续输入y")
		if quit == "x" || quit == "X" {
			break
		}
	}
}
