package rediswrapper

import (
	"errors"
	"runtime"
	"search_engine/internal/util/log"
	"search_engine/internal/util/tools"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

type redisWrapper struct {
	connPool *redis.Pool
}

func NewRedis(host string, port int, password string, index int, timeout time.Duration) (*redisWrapper, error) {
	rd := new(redisWrapper)
	pool := newPool(host, port, password, index, timeout)
	rd.connPool = pool
	_, err := rd.connPool.Dial()
	return rd, err
}

func newPool(host string, port int, password string, index int, timeout time.Duration) *redis.Pool {
	pool := &redis.Pool{
		MaxActive:   3 * runtime.NumCPU(),
		MaxIdle:     2 * runtime.NumCPU(),
		IdleTimeout: 1 * time.Millisecond,
		Dial: func() (redis.Conn, error) {
			if conn, err := redis.Dial("tcp", host+":"+strconv.Itoa(port),
				redis.DialPassword(password), redis.DialDatabase(index),
				redis.DialReadTimeout(timeout*time.Millisecond),
				redis.DialWriteTimeout(timeout*time.Millisecond),
				redis.DialConnectTimeout(timeout*time.Millisecond)); err != nil {
				return nil, err
			} else {
				return conn, nil
			}
		},
	}
	return pool
}

func (rd *redisWrapper) Set(k []byte, v []byte) error {
	defer func(cost func() time.Duration) {
		log.Warnf("cost: %.3f ms", float64(cost().Microseconds())/1000.0)
	}(tools.TimeCost())

	log.Debugf("key:%s", tools.Bytes2Str(k))
	conn := rd.connPool.Get()
	defer conn.Close()
	_, err := conn.Do("Set", k, v)
	return err
}

func (rd *redisWrapper) Get(k []byte) ([]byte, error) {
	defer func(cost func() time.Duration) {
		log.Warnf("cost: %.3f ms", float64(cost().Microseconds())/1000.0)
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

func (rd *redisWrapper) Delete(k []byte) error {
	conn := rd.connPool.Get()
	defer conn.Close()
	_, err := conn.Do("Del", k)
	return err
}

func (rd *redisWrapper) Close() error {
	rd.connPool.Close()
	return nil
}
