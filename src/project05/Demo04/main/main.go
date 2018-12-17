package main

import (
	"fmt"
	"dds/errors"
)

//错误机制
func main() {
	test02()
	test03()
	fmt.Println("main 代码")
}

func test03()  {
	defer func() {
		err :=recover()//recover 可以捕获到异常
		if err!=nil{
			fmt.Println(err)
		}
	}()
	num1:=10
	num2 :=0
	s:=num1/num2
	fmt.Println(s)
}
func readConf(name string) (err error)  {
	if name =="config.ini"{
		return nil
	}else {
		return errors.RawNew("ww","myerror")
	}
}

func test02()  {
	err:=readConf("ocnfig.ini")
	if err !=nil{
		panic(err)
	}
}