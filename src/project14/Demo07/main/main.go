package main

import (
	"fmt"
	"time"
)

func putNum (intChan chan int)  {
	for i:=0;i<=80000;i++{
		intChan <-i
	//	fmt.Println("write",i)
	}
	close(intChan)
}

func primeNum (intChan chan int, primeChan chan int,exitChan chan bool)  {
//	var num int
	var flag bool
	for {
		num ,ok:=<-intChan
		if !ok{
			break
		}
		flag =true
		for i:=2;i<num ;i++{
			if num%i==0{
				flag=false
				break
			}
		}
		if flag{
			primeChan <-num
		}
	}
	fmt.Println("有一个协程完成工作")
	exitChan<-true
}

func main() {
	intChan :=make(chan int,1000)
	primeChan :=make(chan int ,20000)
	exitChan :=make(chan bool,4)
	start:=time.Now().UnixNano()
	go putNum(intChan)
	for i:=0;i<4;i++{
		go primeNum(intChan,primeChan,exitChan)
	}

	go func(){
		for i:=0;i<4;i++{
			<-exitChan
		}
		close(primeChan)
	}()
	end:=time.Now().UnixNano()
	fmt.Println("time:",end-start)
	for{
		res,ok :=<-primeChan
		if !ok{
			break
		}
		fmt.Println(res,"\n")
	}
	fmt.Println("主线程退出")
	}



