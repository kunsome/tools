package file_helper

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
)

// WriteCsv 二维数组输出csv文件
func WriteCsv(fileName string, lines [][]string) {
	dir := filepath.Dir(fileName)
	if !IsFileExist(dir) {
		CreateDir(dir)
	}
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// 写入UTF-8 BOM，避免使用Microsoft Excel打开乱码
	file.WriteString("\xEF\xBB\xBF")
	// 写入csv行

	writer := csv.NewWriter(file)
	writer.Comma = ','
	for _, v := range lines {
		writer.Write(v)
	}

	writer.Flush()
}
