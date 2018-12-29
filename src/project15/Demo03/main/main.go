package main

import (
	"reflect"
	"fmt"
)

func reflect01(b interface{})  {
	rVal:=reflect.ValueOf(b)
	fmt.Println(rVal.Kind())
	rVal.Elem().SetInt(32)
	//rVal.SetInt(32)   # error
}

func main() {

	var num int =10
	reflect01(&num)
	fmt.Println(num)


}