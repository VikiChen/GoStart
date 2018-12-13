package main

import "fmt"


var K="2"
var name  =	"first string"

var(
	K2="2"
	name2  =	"first string"
)

func main() {
	var i =10            //直接赋值
	var name  =	"this is string"
	fmt.Println(name)
	fmt.Println(i)
	//不赋值 取默认
	var t int           //声明类型方式
	fmt.Println(t)
	var x string
	fmt.Println(x)
	var y float64
	fmt.Println(y)
	var z float64
	fmt.Println(z)
	fmt.Println("--------------")
	num :="l am string"   // := 形式
	fmt.Println(num)

	//多变量声明
	var v1,v2,v3 int
	fmt.Println(v1,v2,v3)
	var v4,v5,v6 =100,"tom",20
	fmt.Println(v4,v5,v6)
	v7,v8,v9 :=100,"tom",20
	fmt.Println(v7,v8,v9)
	fmt.Println("--------------")
	fmt.Println(K)
	fmt.Println(name)  //输出的是局部变量
}
