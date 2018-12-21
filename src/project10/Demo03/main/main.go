package main

import "fmt"

type Person struct {
	Name string
	Age int
}

type Student struct {
	Person
	Score int
}


//继承
func main() {
	ppp := Student{}
	ppp.Name="22"
	ppp.Age=2
	ppp.Score=2
	fmt.Println(ppp)
}
