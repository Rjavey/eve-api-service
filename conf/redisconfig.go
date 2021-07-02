package conf

import (
	"eve-api-service/log"
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"reflect"
	"time"
)

var RedisPool *redigo.Pool

func connRedis(addr string, password string) {
	RedisPool = PoolInitRedis(addr, password)
}

func PoolInitRedis(server string, password string) *redigo.Pool {
	return &redigo.Pool{
		MaxIdle:     5, //空闲数
		IdleTimeout: 5 * 60 * 60 * time.Second,
		MaxActive:   10, //最大数
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func Set(k string, v string) {
	_, err := RedisPool.Get().Do("SET", k, v)
	if err != nil {
		log.Error.Printf("redis set err:%v", err)
	}
}

func Get(k string) string {
	v, err := redigo.String(RedisPool.Get().Do("GET", k))
	if err != nil {
		log.Error.Printf("redis set err:%v", err)
	}
	return v
}

func Del(k string) {
	_, err := RedisPool.Get().Do("DEL", k)
	if err != nil {
		log.Error.Printf("redis del err:%v", err)
	}
}

func Setex(k string, v string, ex string) {
	_, err := RedisPool.Get().Do("SET", k, v, "EX", ex)
	if err != nil {
		log.Error.Printf("redis setex err:%v", err)
	}
}

func Expire(k string, ex string) {
	n, err := RedisPool.Get().Do("EXPIRE", k, ex)
	if err != nil {
		log.Error.Printf("redis expire err:%v", err)
	} else if n != int64(1) {
		log.Error.Printf("redis del err:%v", err)
	}
}

func Exist(k string) bool {
	exist, err := redigo.Bool(RedisPool.Get().Do("EXISTS", k))
	if err != nil {
		fmt.Println("err while checking keys:", err)
	}
	return exist
}

func Incr(k string) {
	_, err := RedisPool.Get().Do("INCR", k)
	if err != nil {
		fmt.Println("err while incr key:", err)
	}
}

func Setnx(k string, v string) bool {
	exist, err := redigo.Bool(RedisPool.Get().Do("SETNX", k, v))
	if err != nil {
		fmt.Println("err while checking keys:", err)
	}
	return exist
}

func Scan(k string) []string {
	iter := 0
	var keys []string
	for {
		if arr, err := redigo.MultiBulk(RedisPool.Get().Do("SCAN", iter, "MATCH", "*"+k+"*")); err != nil {
			log.Error.Printf("scan keys err:%v", err)
		} else {
			iter, _ = redigo.Int(arr[0], nil)
			key, _ := redigo.Strings(arr[1], nil)
			for _, value := range key {
				keys = append(keys, value)
			}
		}
		if iter == 0 {
			break
		}
	}
	return keys
}

func Lpush(k string, v string) {
	_, err := RedisPool.Get().Do("LPUSH", k, v)
	if err != nil {
		log.Error.Printf("redis lpush err:%v", err)
	}
}

func Lrange(k string, s1 int64, s2 int64) []string {
	vals, err := redigo.Values(RedisPool.Get().Do("RPUSH ", k, s1, s2))
	if err != nil {
		log.Error.Printf("redis rpush err:%v", err)
	}

	var res []string
	for _, v := range vals {
		if reflect.TypeOf(v).String() == "[]uint8" {
			str := string(v.([]byte))
			res = append(res, str)
		}
	}
	return res
}

func Zset(k string, v string, score int64) {
	_, err := RedisPool.Get().Do("ZADD", k, score, v)
	if err != nil {
		log.Error.Printf("redis zadd err:%v", err)
	}
}

func ZrangeByScore(k string, s1 int64, s2 int64) []string {
	vals, err := redigo.Values(RedisPool.Get().Do("zrange", k, s1, s2))
	if err != nil {
		log.Error.Printf("redis zadd err:%v", err)
	}

	var res []string
	for _, v := range vals {
		if reflect.TypeOf(v).String() == "[]uint8" {
			str := string(v.([]byte))
			res = append(res, str)
		}
	}
	return res
}
