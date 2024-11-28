// redis初始化连接

package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-vue/common/config"
)

var (
	RedisDb *redis.Client
)

// 1.创建Redis客户端连接 2.测试连接是否成功 3.返回连接初始化的结果。

func SetupRedisDb() error {
	var ctx = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Address,
		Password: config.Config.Redis.Password,
		DB:       0,
	})
	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	return nil
}
