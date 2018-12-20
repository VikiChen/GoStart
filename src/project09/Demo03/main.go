package main

import "fmt"

type stu struct {
	Name string
	Age int
}

type stu1 struct {
	Name string
	Age int
}

//结构体
func main() {
	p1:=stu{"s",2}
	p2:=stu1{"ss",2}
	fmt.Println(p1)
	fmt.Println(p2)

	p2=stu1(p1)

}
