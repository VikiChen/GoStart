package main

import (
	"time"
	"fmt"
)

func main() {

	fmt.Println(time.Now())
	now :=time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Format("01"))
	time.Sleep(time.Millisecond*10)
}
