package file_helper

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// UrlDownload 下载url文件到本地
func UrlDownload(url string, path string) (err error) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	dir := filepath.Dir(path)
	if !IsFileExist(dir) {
		err = CreateDir(dir)
		if err != nil {
			return
		}
	}
	// Create output file
	out, err := os.Create(path)
	if err != nil {
		return
	}
	defer out.Close()
	// copy stream
	_, err = io.Copy(out, resp.Body)
	return
}
