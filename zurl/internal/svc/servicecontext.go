package svc

import (
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zlink/zurl/internal/config"
	"zlink/zurl/model"
	"zlink/zurl/sequence"
)

type ServiceContext struct {
	Config config.Config
	// 短链接存储model
	ZUrlModel model.ShortUrlMapModel
	// 短链接序列号model
	SeqModel model.SequenceModel

	// Mysql 自定义的数据库key查询
	Sequence sequence.SequenceSql
	// Redis 自定义的数据库key查询 二两选一个
	//RdsSequence sequence.SequenceSql

	// 布隆过滤器
	// bloom filter
	Filter *bloom.Filter
}

func NewServiceContext(c config.Config) *ServiceContext {
	connZurlmysql := sqlx.NewMysql(c.ZUrlMysql.DataSource)
	connSeqmysql := sqlx.NewMysql(c.SequenceMysql.DataSource)

	store := redis.MustNewRedis(c.CacheRedis[0].RedisConf)
	// 新建布隆过滤器
	//store, _ := redis.NewRedis(c.BloomRedisHost)
	filter := bloom.New(store, "bloom_Url", 1024)

	return &ServiceContext{
		Config:    c,
		ZUrlModel: model.NewShortUrlMapModel(connZurlmysql, c.CacheRedis),
		SeqModel:  model.NewSequenceModel(connSeqmysql, c.CacheRedis),
		Sequence:  sequence.NewMysql(c.SequenceMysql.DataSource),
		//RdsSequence: sequence.NewRedis(c.RedisUrl.Host),
		Filter: filter,
	}
}
