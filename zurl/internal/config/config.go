package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	// 缓存
	CacheRedis cache.CacheConf
	// 短链接存储数据库
	ZUrlMysql struct {
		DataSource string
	}
	// 短链接序列号锚点数据库
	SequenceMysql struct {
		DataSource string
	}
	// mysql 或者 redis
	RedisUrl struct {
		Host string
	}
	Base62Chars string

	ZUrlDoamin string
}
