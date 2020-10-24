package redis

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"

	"github.com/Antony15/goodWorkLabs-Test/constants"

	"github.com/Antony15/goodWorkLabs-Test/logger"
)

var (
	client = &redisClient{}
)

type redisClient struct {
	c *redis.Client
}

//Initialize get the redis client
func Initialize() *redisClient {
	c := redis.NewClient(&redis.Options{
		Addr:       constants.RedisHost + ":" + constants.RedisPort,
		Password:   constants.RedisPass,
		DB:         constants.RedisDbName,
		MaxConnAge: 1,
	})
	//defer c.Close()
	if err := c.Ping().Err(); err != nil {
		logger.Log.Println("Unable to connect to redis " + err.Error())
		panic(err)
	}
	client.c = c
	return client
}

//GetKey get's the value of a given key
func (client *redisClient) GetKey(key string, src interface{}) error {
	val, err := client.c.Get(key).Result()
	if err == redis.Nil || err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), &src)
	if err != nil {
		return err
	}
	return nil
}

//SetKey set's the value of a given key
func (client *redisClient) SetKey(key string, value interface{}, expiration time.Duration) error {
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = client.c.Set(key, cacheEntry, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
