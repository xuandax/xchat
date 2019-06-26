package datasource

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"sync"
	"time"
)

var (
	instanceRedis *redisConn
	mu            sync.Mutex
)

type redisConn struct {
	pool      *redis.Pool
	showDebug bool
}

func (r *redisConn) Do(commandName string, args ...interface{}) (replay interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	t1 := time.Now().UnixNano()
	replay, err = conn.Do(commandName, args...)
	if err != nil {
		e := conn.Err()
		if e != nil {
			log.Println("datasource redis Do err = ", err, e)
		}
	}
	t2 := time.Now().UnixNano()
	if r.showDebug {
		log.Printf("[redis] [info] [time:%dus] command=%s, args=%s, replay=%s, err=%s\n", (t2-t1)/1000, commandName, args, replay, err)
	}
	return
}

func InstanceRedis() *redisConn {
	if instanceRedis != nil {
		return instanceRedis
	}
	mu.Lock()
	defer mu.Unlock()
	if instanceRedis != nil {
		return instanceRedis
	}
	return newRedisConn()
}

func (r *redisConn) ShowDebug(b bool) {
	r.showDebug = b
}

func newRedisConn() *redisConn {
	host := "120.79.167.42"
	port := 3306
	pool := redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
			if err != nil {
				log.Fatal("datasource redis newRedisConn err = ", err)
				return nil, err
			}
			return conn, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         10000, //最大连接数
		MaxActive:       10000, //最大活跃数
		IdleTimeout:     0,
		Wait:            false,
		MaxConnLifetime: 0,
	}
	instanceRedis = &redisConn{
		pool:      &pool,
		showDebug: false,
	}
	return instanceRedis
}
