package main

import (
	"strconv"
	"fmt"
)

//数据类型与String类型 相互转化
func main() {

	var str string="true"
	var b bool
	b,_=strconv.ParseBool(str)
	fmt.Printf("%T,%v\n",b,b)

	var str2 string="12345"
	var n1 int64
	n1,_=strconv.ParseInt(str2,10,64)
	fmt.Printf("%T,%v\n",n1,n1)

	var str3 string="12345.222"
	var n2 float64
	n2,_=strconv.ParseFloat(str3,64)
	fmt.Printf("%T,%v\n",n2,n2)
	}
