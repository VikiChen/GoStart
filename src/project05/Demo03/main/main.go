package main

import (
	"strconv"
	"time"
	"fmt"
)

func main() {
	start :=time.Now().Unix()
	test03()
	end :=time.Now().Unix()
	fmt.Println(end-start)
}

func test03()  {
	str :=""
	for i:=0;i<100000;i++{
		str +="hello" +strconv.Itoa(i)
	}

}