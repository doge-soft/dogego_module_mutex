package dogego_module_mutex

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type RedisMutex struct {
	RedisClient *redis.Client
}

func NewRedisMutex(redis_Client *redis.Client) *RedisMutex {
	return &RedisMutex{
		RedisClient:redis_Client,
	}
}

func (mutex *RedisMutex) Lock(lock_name string, lock_time time.Duration) bool {
	result, err := mutex.RedisClient.SetNX(
		fmt.Sprintf("lock:%s", lock_name), "true", lock_time).Result()

	if err != nil {
		return false
	}

	return result
}

func (mutex *RedisMutex) UnLock(lock_name string) error {
	return mutex.RedisClient.Del(fmt.Sprintf("lock:%s", lock_name)).Err()
}
