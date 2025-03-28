package svc

import (
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
}

func NewServiceContext(c config.Config) *ServiceContext {
	connZurlmysql := sqlx.NewMysql(c.ZUrlMysql.DataSource)
	connSeqmysql := sqlx.NewMysql(c.SequenceMysql.DataSource)
	return &ServiceContext{
		Config:    c,
		ZUrlModel: model.NewShortUrlMapModel(connZurlmysql, c.CacheRedis),
		SeqModel:  model.NewSequenceModel(connSeqmysql, c.CacheRedis),
		Sequence:  sequence.NewMysql(c.SequenceMysql.DataSource),
		//RdsSequence: sequence.NewRedis(c.RedisUrl.Host),
	}
}
