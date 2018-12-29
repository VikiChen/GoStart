package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

var pool *redis.Pool

func init() {
	pool=&redis.Pool{
		MaxIdle:8,
		MaxActive:0,
		IdleTimeout:100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp","localhost:6379")
		},
	}
}

func main() {
	conn := pool.Get()
	fmt.Println(conn)
	defer conn.Close()
	_, err := conn.Do("Set", "name2", "tom2")
	if err!= nil{
		fmt.Println(err)
		return
	}

}