package datastore

import (
	"fmt"
	"globetrotter_api/config"
	"log"

	"github.com/redis/go-redis/v9"
)

func RedisClient() *redis.Client {

	redisClient := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%v:%v", config.C.Redis.Host, config.C.Redis.Port),
		Password:    config.C.Redis.Password,
		DB:          config.C.Redis.Db,
		PoolFIFO:    true,
		MaxRetries:  2,
		ReadTimeout: -1,
	})

	log.Println("Connected to redis")

	return redisClient
}
