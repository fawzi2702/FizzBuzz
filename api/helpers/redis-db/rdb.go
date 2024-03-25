package redisdb

import (
	"context"
	"fmt"

	"github.com/fawzi2702/FizzBuzz/api/helpers/environment"
	"github.com/redis/go-redis/v9"
)

var Context = context.Background()

var Client *redis.Client

func SetupRedisClient() error {
	if Client != nil {
		return nil
	}

	addr, err := environment.Get("REDIS_ADDR")
	if err != nil {
		return err
	}

	dbIndex, err := environment.GetInt("REDIS_DB_INDEX")
	if err != nil {
		return err
	}

	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       dbIndex,
	})

	if _, err := c.Ping(Context).Result(); err != nil {
		return err
	} else {
		fmt.Println("RedisDB connected")
	}

	Client = c

	return nil
}
