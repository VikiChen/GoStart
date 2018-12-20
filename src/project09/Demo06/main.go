package main

import (
	"fmt"
)

type MethodUtils struct {

}

func (methodUtils MethodUtils)test()  {
	fmt.Println("1")
}

//结构体
func main() {
	p1:=MethodUtils{}
	p1.test()


}
