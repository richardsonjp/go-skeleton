package file

import (
	"net/http"
	"os"
)

// Exists check file exist
func Exists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !(info.IsDir())
}

// Get remote file via HTTP request and then turn it into bytes
func GetRemoteFileInBytes(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	content := make([]byte, resp.ContentLength)
	_, err = resp.Body.Read(content)
	defer resp.Body.Close()

	return content, nil
}
