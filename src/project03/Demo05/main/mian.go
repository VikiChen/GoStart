package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println(time.Now().Unix())
	for{
		//rand.Seed(time.Now().UnixNano())
		var x =rand.Intn(100)+1
		fmt.Println(x)
		if x==99 {
			break
		}
	}

}
