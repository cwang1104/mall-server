package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var RedisConn redis.Conn
var Err error

func init() {
	RedisConn, Err = redis.Dial("tcp", "127.0.0.1:6379")
	if Err != nil {
		fmt.Println(Err)
		panic(Err)
	}
}
