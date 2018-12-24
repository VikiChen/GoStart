package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string `json:"name"`
	Age int
	Birthday string
	Sal float64
	Skill string
}

func testStuct()  {
	var moster Monster
	moster.Name="牛"
	moster.Age=2
	moster.Birthday="20191-21-21"
	moster.Sal=2020
	moster.Skill="jhhh"

	bytes, _ := json.Marshal(moster)
	fmt.Println(string(bytes))
}

func testMap()  {
	var a map[string]interface{}
	//使用map 先make
	a=make(map[string]interface{})
	a["name"]="红孩儿"
	a["age"]="30"
	bytes, _ := json.Marshal(a)
	fmt.Println(string(bytes))
}

func testSlice()  {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1=make(map[string]interface{})
	m1["name"]="红孩儿"
	m1["age"]="30"
	slice = append(slice, m1)
	var m2 map[string]interface{}
	m2=make(map[string]interface{})
	m2["name"]="红孩儿2"
	m2["age"]="302"
	slice = append(slice, m2)
	bytes, _ := json.Marshal(slice)
	fmt.Println(string(bytes))
}
func testFloat64(){
	var byn float64=23.355
	bytes, _ := json.Marshal(byn)
	fmt.Println(string(bytes))
	}

func main() {

testStuct()
testMap()
testSlice()
testFloat64()
}

