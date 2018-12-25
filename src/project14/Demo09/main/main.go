package main

import (
	"time"
	"fmt"
)

func sayHello()  {
	for i:=0;i<10;i++{
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}

func  test()  {
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println("异常",err)
		}
	}()
	var myMap map[int]string
	myMap[0]="golang"
}

func main() {
	go sayHello()
	go test()

	for i:=0;i<10;i++{
		fmt.Println("main",i)
		time.Sleep(time.Second)
	}
	}



