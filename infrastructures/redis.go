package infrastructures

import (
	"context"
	"fmt"
	"pi/config"

	"github.com/redis/go-redis/v9"
)

func NewRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Get().Redis.Host, config.Get().Redis.Port),
		Password: config.Get().Redis.Password,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		panic("error connect redis client: " + err.Error())
	}

	return client
}
