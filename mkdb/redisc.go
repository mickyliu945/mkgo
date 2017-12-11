package mkdb

import (
	"github.com/garyburd/redigo/redis"
	"mkgo/mkconfig"
	"time"
)

var redisPool *redis.Pool

//func getRedisConn() (c redis.Conn) {
//	c, err := redis.Dial("tcp", mkconfig.Config.MKGo.Redis.Host)
//	if err != nil {
//		mklog.Logger.Error("[redisc]", zap.Error(err))
//		return
//	}
//	return c
//}

func getPool() *redis.Pool {
	if redisPool == nil {
		redisPool = &redis.Pool{
			MaxIdle:    mkconfig.Config.MKGo.Redis.MaxIdle,
			MaxActive:   mkconfig.Config.MKGo.Redis.MaxActive,
			Wait:        true,
			IdleTimeout: 120 * time.Second,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", mkconfig.Config.MKGo.Redis.Host)
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < 10 * time.Second {
					return nil
				}
				_, err := c.Do("PING")
				return err
			},
		}
	}
	return redisPool
}

func RedisSET(key string, value string) (string, error) {
	conn := getPool().Get()
	defer  conn.Close()
	reply, err := redis.String(conn.Do("SET", key, value))
	return reply, err
}

func RedisSETX(key string, value string, seconds int) (string, error) {
	conn := getPool().Get()
	defer  conn.Close()
	reply, err := redis.String(conn.Do("SETEX", key, seconds, value))
	return reply, err
}

func RedisGET(key string) (string, error) {
	conn := getPool().Get()
	defer conn.Close()
	reply, err := redis.String(conn.Do("GET", key))
	return reply, err
}

func RedisDEL(key string) (string, error)  {
	conn := getPool().Get()
	defer conn.Close()
	reply, err := redis.String(conn.Do("DEL", key))
	return reply, err
}
