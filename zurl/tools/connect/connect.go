package connect

import (
	"net/http"
)

// CheckWebsite 定义一个函数，用来判断传入的网址是否通
func CheckWebsite(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true
	}

	return false
}
