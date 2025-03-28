package main

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func main() {
	conf := redis.RedisConf{
		Host:        "127.0.0.1:6379",
		Type:        "node",
		Pass:        "",
		Tls:         false,
		NonBlock:    false,
		PingTimeout: time.Second,
	}
	rds := redis.MustNewRedis(conf)
	ctx := context.Background()
	err := rds.SetCtx(ctx, "key", "hello world")

	if err != nil {
		logc.Error(ctx, err)
	}

	v, err := rds.GetCtx(ctx, "key")
	if err != nil {
		logc.Error(ctx, err)
	}
	fmt.Println(v)
}
