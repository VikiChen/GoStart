package main

import (
	"fmt"
)



func main() {
	intChan :=make(chan int,10)
	for i:=0;i<10;i++{
		intChan <-i
	}
	stringChan :=make(chan string,5)
	for i:=0;i<5;i++{
		stringChan <-"hello"+fmt.Sprintf("%d",i)
	}
	label:
	for {
		select {
		case v:=<-intChan:
			fmt.Println(v)
		case v:=<-stringChan:
			fmt.Println(v)
		default:
			fmt.Println("null")
		    break label
		}
	}

	}



