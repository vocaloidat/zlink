package main

import (
	"flag"
	"fmt"
	"zlink/zurl/internal/config"
	"zlink/zurl/internal/handler"
	"zlink/zurl/internal/svc"
	"zlink/zurl/tools/base62"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/zurl-api.yaml", "the config file")
var blackListConfigFile = flag.String("b", "etc/UrlBlackList.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	// 初始化自己的base62
	base62.NewBase62Tool(c.Base62Chars)

	// 初始化黑名单
	conf.MustLoad(*blackListConfigFile, &config.BlackLC)
	config.BlackLC.ToMap() // 将黑名单转为map,这样校验更快

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
