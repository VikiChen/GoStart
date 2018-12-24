package main

import (
	"encoding/json"
	"fmt"
)


//{"name":"牛","Age":2,"Birthday":"20191-21-21","Sal":2020,"Skill":"jhhh"}
//{"age":"30","name":"红孩儿"}
//[{"age":"30","name":"红孩儿"},{"age":"302","name":"红孩儿2"}]

type Monster struct {
	Name string `json:"name"`
	Age int
	Birthday string
	Sal float64
	Skill string
}

func unmarshalStruct()  {
	str:="{\"name\":\"牛\",\"Age\":2,\"Birthday\":\"20191-21-21\",\"Sal\":2020,\"Skill\":\"jhhh\"}"
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err!=nil{
		fmt.Println("unmarshal err=",err)
	}
	fmt.Println(monster)

}

func unmarshalMap()  {
	str:="{\"age\":\"30\",\"name\":\"红孩儿\"}"
	var a map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err!=nil{
		fmt.Println("unmarshal err=",err)
	}
	fmt.Println(a)

}


func unmarshalSlice()  {
	str:="[{\"age\":\"30\",\"name\":\"红孩儿\"},{\"age\":\"302\",\"name\":\"红孩儿2\"}]"
	var a []map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err!=nil{
		fmt.Println("unmarshal err=",err)
	}
	fmt.Println(a)

}

func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}

