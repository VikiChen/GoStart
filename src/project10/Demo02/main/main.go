package main

import (
	"project10/Demo02/model"
	"fmt"
)

//抽象
func main() {

	person1:=model.NewPerson("test")
	person1.SetAge(10)
	person1.SetSal(3222)
	fmt.Println(person1)
}
