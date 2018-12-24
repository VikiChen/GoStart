package main

import (
	"runtime"
	"fmt"
)

func main() {
	cpu := runtime.NumCPU()
	fmt.Println(cpu)
	runtime.GOMAXPROCS(2)
}



