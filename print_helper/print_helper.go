package print_helper

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// PrintStruct 格式化为json后,打印对象
// 参数：
//   - v 被打印对象
func PrintStruct(v interface{}) {
	b, _ := json.Marshal(v)
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")
	if err != nil {
		fmt.Println("打印时发生错误：", err)
	}
	fmt.Println(out.String())
}
