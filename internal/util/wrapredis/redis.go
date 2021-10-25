package wrapredis

import (
	"errors"
	"runtime"
	"search_engine/internal/util/log"
	"search_engine/internal/util/tools"
	"time"

	"github.com/gomodule/redigo/redis"
)

type wrapRedis struct {
	connPool *redis.Pool
}

func NewRedis(host string, port string, password string, index int, timeout int) (*wrapRedis, error) {
	rd := new(wrapRedis)
	pool := newPool(host, port, password, index, timeout)
	rd.connPool = pool
	return rd, nil
}

func newPool(host string, port string, password string, index int, timeout int) *redis.Pool {
	return &redis.Pool{
		MaxActive:   3 * runtime.NumCPU(),
		MaxIdle:     2 * runtime.NumCPU(),
		IdleTimeout: 1 * time.Millisecond,
		Dial: func() (redis.Conn, error) {
			if conn, err := redis.Dial("tcp", host+":"+port,
				redis.DialPassword(password), redis.DialDatabase(index),
				redis.DialReadTimeout(time.Duration(timeout)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(timeout)*time.Millisecond),
				redis.DialConnectTimeout(time.Duration(timeout)*time.Millisecond)); err != nil {
				return nil, err
			} else {
				return conn, nil
			}
		},
	}
}

func (rd *wrapRedis) Set(k []byte, v []byte) error {
	defer func(cost func() time.Duration) {
		log.Warnf("trackid:%d, cost: %.3f ms", 0, float64(cost().Microseconds())/1000.0)
	}(tools.TimeCost())

	log.Debugf("key:%s", tools.Bytes2Str(k))
	conn := rd.connPool.Get()
	defer conn.Close()
	_, err := conn.Do("Set", k, v)
	return err
}

func (rd *wrapRedis) Get(k []byte) ([]byte, error) {
	defer func(cost func() time.Duration) {
		log.Warnf("trackid:%d, cost: %.3f ms", 0, float64(cost().Microseconds())/1000.0)
	}(tools.TimeCost())

	log.Debugf("key:%s", tools.Bytes2Str(k))
	conn := rd.connPool.Get()
	defer conn.Close()
	v, err := redis.Bytes(conn.Do("Get", k))
	if err != nil || len(v) == 0 {
		return nil, errors.New("data empty")
	}
	return v, err
}

func (rd *wrapRedis) Delete(k []byte) error {
	conn := rd.connPool.Get()
	defer conn.Close()
	_, err := conn.Do("Del", k)
	return err
}

func (rd *wrapRedis) Close() error {
	rd.connPool.Close()
	return nil
}
