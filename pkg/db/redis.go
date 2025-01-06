// By Saif Hamdan, Team Lead
// Date: 2025/1/6
package db

import (
	"context"
	"fmt"

	redisPkg "github.com/redis/go-redis/v9"
)

type redis struct {
	*redisPkg.Client
}

func (d *DB) newRedis() (*redis, error) {
	addr := fmt.Sprintf("%s:%s", d.cfg.RedisHost, d.cfg.RedisPort)

	d.logger.Infof("connecting to  redis: %s", addr)

	rdb := redisPkg.NewClient(&redisPkg.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Perform a health check
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	d.logger.Infof("redis connection established %s", addr)

	return &redis{Client: rdb}, nil
}
