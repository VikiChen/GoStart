package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{})  {

}

type Monster struct {
	Name string `json:"name"`
	Age int `json:"monster_age"`
	Socre float32
	Sex string
}

func (s Monster) Print()  {
	fmt.Println("----start-----")
	fmt.Println(s)
	fmt.Println("-----end------")
}

func (s Monster) GetSum(n1,n2 int) int {
	return n1+n2
}

func (s Monster) Set(name string,age int,score float32,sex string)  {
	s.Name=name
	s.Age=age
	s.Socre=score
	s.Sex=sex
}

func TestStruct(a interface{})  {
	typ :=reflect.TypeOf(a)
	val :=reflect.ValueOf(a)
	kd :=val.Kind()
	if kd!=reflect.Struct{
		fmt.Println("expect struct")
		return
	}

	num :=val.NumField()
	fmt.Printf("struct has %d ,fields\n",num)
	for i:=0;i<num ;i++  {
		fmt.Printf("field %d :vale=%v\n",i,val.Field(i))
		tagVal:=typ.Field(i).Tag.Get("json")
		if tagVal !=""{
			fmt.Printf("field %d :vale=%v\n",i,tagVal)
		}

	}

	numOfMethod :=val.NumMethod()
	fmt.Printf("has %d methods\n",numOfMethod)

	val.Method(1).Call(nil)

	var params []reflect.Value
	params =append(params,reflect.ValueOf(10))
	params=append(params,reflect.ValueOf(40))
	res:=val.Method(0).Call(params)
	fmt.Println(res[0].Int())

}


func main() {

	var a Monster=Monster{
		Name:"Niu",
		Age:22,
		Socre:20.3,
	}

	TestStruct(a)
	


}