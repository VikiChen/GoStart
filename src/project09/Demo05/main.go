package main

import (
	"fmt"
)

type stu struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func (s stu)test() int  {
	return s.Age
}

//结构体
func main() {
	p1:=stu{"s",2}
	fmt.Println(p1.test())


}
