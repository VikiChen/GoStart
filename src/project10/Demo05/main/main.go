package main

import "fmt"

type Ainterface interface {
	Say()
}

type Stu struct {
	Name string
}


func (p Stu) Say()  {
	fmt.Println("Stu say")
}


//接口
func main() {
	var stu Stu
	var a Ainterface =stu
	a.Say()
}
