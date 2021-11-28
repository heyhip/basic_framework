package cache

import (
	"basic_framework/configs"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var Orm *redis.Client

func init() {
	Orm = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", configs.Yaml.Redis.Host, configs.Yaml.Redis.Port),
		Password: configs.Yaml.Redis.Auth,
		DB:       configs.Yaml.Redis.Select,
		Network:  configs.Yaml.Redis.Network,
		PoolSize: configs.Yaml.Redis.PoolSize,
	})
}

// 加锁
func Lock(key string, value interface{}, second int) bool {
	return Orm.SetNX(context.TODO(), key, value, time.Duration(second)*time.Second).Val()
}

// 获取锁信息
func GetLock(key string) string {
	return Orm.Get(context.TODO(), key).Val()
}

// 解锁
func UnLock(key string) bool {
	return Orm.Del(context.TODO(), key).Val() > 0
}

// 增加缓存时长
func Expire(key string, second int) bool {
	return Orm.Expire(context.TODO(), key, time.Duration(second)*time.Second).Val()
}

// 进行缓存递增计数
func IncrCacheNumber(key string, second int) (i int64) {
	i = Orm.Incr(context.TODO(), key).Val()
	Orm.Expire(context.TODO(), key, time.Duration(second)*time.Second)
	return
}

// 进行缓存递减计数
func DecrCacheNumber(key string, second int) (i int64) {
	i = Orm.Decr(context.TODO(), key).Val()
	Orm.Expire(context.TODO(), key, time.Duration(second)*time.Second)
	return
}
