package main

import (
	"fmt"
	"encoding/json"
)

type stu struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type stu1 struct {
	Name string
	Age int
}

//结构体
func main() {
	p1:=stu{"s",2}
	fmt.Println(p1)
	jsonStr,err:=json.Marshal(p1)
	fmt.Println(string(jsonStr))
	fmt.Println(err)

}
