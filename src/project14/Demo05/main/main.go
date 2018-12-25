package main

import "fmt"

func main() {
	intChan :=make(chan int ,3)
	intChan <-2
	intChan <-3
	close(intChan)
	<-intChan

	intChan2 :=make(chan int,100)
	for i:=0;i<100;i++{
		intChan2<-i*2
	}

	close(intChan2)
	for v :=range intChan2{
		fmt.Println(v)
	}
}



