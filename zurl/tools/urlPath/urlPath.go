package urlPath

import (
	"net/url"
	"path"
)

// GetBasePath 获取连接路径，传入www.baidu.com/weqew?name=wyc 返回weqew
func GetBasePath(pathStr string) (basePath string, err error) {
	MyUrl, err := url.Parse(pathStr)
	if err != nil {
		return "", err
	}
	basePath = path.Base(MyUrl.Path)
	return basePath, nil
}
