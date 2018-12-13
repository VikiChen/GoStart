package main

import (
	"fmt"
	"strconv"
)

//数据类型与String类型 相互转化
func main() {

	var num1 int =99
	var num2 float64 =23.456
	var b bool =true
	var c1 byte ='b'
	var str string
	//使用第一种方式
	str =fmt.Sprintf("%d",num1)
	fmt.Printf("%T,%q\n",str,str)

	str =fmt.Sprintf("%f",num2)
	fmt.Printf("%T,%q\n",str,str)

	str =fmt.Sprintf("%t",b)
	fmt.Printf("%T,%q\n",str,str)

	str =fmt.Sprintf("%c",c1)
	fmt.Printf("%T,%q\n",str,str)

	//第二种方式 strconv包
	str =strconv.FormatInt(int64(num1),10)  //第二个参数是 进制
	fmt.Printf("%T,%q\n",str,str)
	str =strconv.FormatFloat(num2,'f',10,64)  //10 小数点10位   64 表示float64
	fmt.Printf("%T,%q\n",str,str)
	str =strconv.FormatBool(b)
	fmt.Printf("%T,%q\n",str,str)


	//Itoa

	var num int =4567
	str =strconv.Itoa(num)
	fmt.Printf("%T,%q\n",str,str)
	}
