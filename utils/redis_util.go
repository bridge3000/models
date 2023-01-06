package utils

import (
	"github.com/garyburd/redigo/redis"
)

type RedisUtil struct {
}

func (this *RedisUtil) GetConn(host string, port string, pwd string, dbName string) (redis.Conn, error) {
	c, err := redis.Dial("tcp", host+":"+port)
	//	c, err := redis.Dial("tcp", host+":7100")
	if err == nil {
		if pwd != "" {
			_, err = c.Do("AUTH", pwd)
		}

		c.Do("SELECT", dbName)
	}

	return c, err
}
