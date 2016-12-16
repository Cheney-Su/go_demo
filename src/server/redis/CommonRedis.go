package redis

import (
	"go/src/server/common"
	"github.com/garyburd/redigo/redis"
	"log"
)

var (
	commonRedis redis.Conn
)

func init() {
	commonRedis = common.Redis
}

func Hmset(key string, myMap map[string]string) bool {
	if ok, err := redis.Bool(commonRedis.Do("HSET", key, "shq", "1111")); ok {
		log.Println(ok)
		return true
	} else {
		log.Print(err)
		return false
	}
}