package main

import (
	"reflect"
	"fmt"
)

//反射
func reflectTest01(b interface{})  {
	rType:=reflect.TypeOf(b)
	fmt.Println(rType)

	rval:=reflect.ValueOf(b)
	fmt.Println(rval)

	n2:=2+rval.Int()
	fmt.Println(n2)

	iv:=rval.Interface()
	num3:=iv.(int)
	fmt.Println(num3)
}

func reflectTest02(b interface{})  {
	rval:=reflect.ValueOf(b)
	iv:=rval.Interface()
	rkind := rval.Kind()
	fmt.Println(rkind)
	fmt.Printf("%T\n",iv)
	num3:=iv.(Student)
	fmt.Println(num3.Name)
}

type Student struct {
	Name string
	Age int
}

func main() {
	var num int =100
	reflectTest01(num)
	stu:=Student{
		Name:"tin",
		Age:2,
	}
	reflectTest02(stu)
	fmt.Println(stu)
}