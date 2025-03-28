package sequence

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	Key = `zurl:maxid`
)

type Redis struct {
	rds *redis.Redis
}

func NewRedis(host string) SequenceSql {
	rds := redis.MustNewRedis(redis.RedisConf{
		Host: host,
	})
	return Redis{
		rds: rds,
	}
}

func (r Redis) Next() (uint64, error) {
	incr, err := r.rds.Incr(Key)
	if err != nil {
		return 0, err
	}
	return uint64(incr), nil
}
