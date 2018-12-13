package main

import "fmt"

//数据类型  string
func main() {

	var add string="北京天安门 110 hello world"
	fmt.Println(add)

	//注意事项
	//一旦赋值 不能变
	var str ="hello"
//	str[0]='a'  //can not change
	fmt.Println(str) //wrong


	//反引号
	str3 :="abc\nabc"
	fmt.Println(str3)
	str4 :=`smfr\nsbdjsd`   //使用反引号
	fmt.Println(str4)

	//拼接字符串
	str4 ="adsd"+"sdsd"+"sddsd"+
		"sd"
	fmt.Println(str4)
}
