package storage

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

type Cache struct {
	RedisPool *redis.Pool
}

func (c *Cache) Get(key string) interface{} {
	client := c.RedisPool.Get()
	defer client.Close()
	data, _ := client.Do("GET", key)
	return data
}

func (c *Cache) Set(key string, val interface{}, timeout time.Duration) error {
	client := c.RedisPool.Get()
	defer client.Close()
	_, err := client.Do("SET", key, val, "EX", timeout.Seconds())
	return err
}
func (c *Cache) IsExist(key string) bool {
	client := c.RedisPool.Get()
	defer client.Close()
	b, _ := redis.Bool(client.Do("EXISTS", key))
	return b
}
func (c *Cache) Delete(key string) error {
	client := c.RedisPool.Get()
	defer client.Close()
	_, err := client.Do("DEL", key)
	return err
}
