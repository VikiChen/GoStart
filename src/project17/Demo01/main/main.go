package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main() {
	conn, e := redis.Dial("tcp", "10.101.12.7:6379")
	if e!=e{
		fmt.Println("连接失败",e)
		return
	}
	defer conn.Close()
	reply, err := conn.Do("AUTH", "123456789")
	if err!=nil{
		fmt.Println("密码错误")
		return
	}
	fmt.Println(reply)
	_, err2 := conn.Do("set", "name", "tomjerrt")
	if err2!=nil{
		fmt.Println("插入失败")
		return
	}
	reply2, err3 :=redis.String(conn.Do("get", "name"))
	if err3!=nil{
		fmt.Println("查询失败")
		return
	}
	fmt.Println(reply2)

}