package wechat

import (
	"github.com/silenceper/wechat/v2/cache"
	"wejh-go/service/redis"
)

func setRedis(wcCache cache.Cache) cache.Cache {
	redisOpts := &cache.RedisOpts{
		Host:        redis.Info.Host + ":" + redis.Info.Port,
		Database:    redis.Info.DB,
		MaxActive:   10,
		MaxIdle:     10,
		IdleTimeout: 60,
	}
	wcCache = cache.NewRedis(redisOpts)
	return wcCache
}
