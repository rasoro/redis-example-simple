package main

import "github.com/go-redis/redis"

var (
	client = &redisClient{}
)

type redisClient struct {
	c *redis.Client
}

func NewRedisClient() *redisClient {
	c := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if err := c.Ping().Err(); err != nil {
		panic("Unable to connect to redis " + err.Error())
	}

	client.c = c

	return client
}
