package main

import "fmt"

func main() {
	var intChan chan int
	intChan =make(chan int,5)
	fmt.Println(intChan)
	intChan <- 10
	num :=22
	intChan <-num
	fmt.Println(len(intChan))
	fmt.Println(cap(intChan))

	num2 :=<-intChan
	fmt.Println(num2)

}



