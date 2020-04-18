/**
*
*@author 吴昊轩
*@create 2020-04-1517:28
 */
package db

import (
	"github.com/go-redis/redis"
	"golang-blog/config"
	"time"
)

var G_redisClient *redis.Client

func InitRedisClient() (err error) {
	cli := redis.NewClient(&redis.Options{
		Addr:        config.G_cfg.RedisInfo.Host,
		Password:    "",
		DB:          config.G_cfg.RedisInfo.Db,
		DialTimeout: 10 * time.Second,
		ReadTimeout: 3 * time.Second,
		PoolSize:    5,
		PoolTimeout: 10 * time.Second,
	})
	_, err = cli.Ping().Result()
	if err != nil {

		return
	}

	G_redisClient = cli
	return
}
