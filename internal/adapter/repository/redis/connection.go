package redis

import (
	"better_auth/internal/config"
	"context"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"
)


var Client *goredis.Client

func InitialRedisConnection (cfg *config.AppConfig) (error) {
	var err error

	Client = goredis.NewClient(&goredis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisConfig.RedisHost, cfg.RedisConfig.RedisPort),
		Password: cfg.RedisConfig.RedisPassword,
		DB:       cfg.RedisConfig.RedisDB,
	})


	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()


	_, err = Client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("error ping redis client")
	}

	return nil
}
