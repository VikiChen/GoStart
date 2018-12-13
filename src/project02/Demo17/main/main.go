package main

import "fmt"

func main() {
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("快输入")
	fmt.Scanln(&name)
	fmt.Println(name)
	fmt.Println("快输入")
	fmt.Scanln(&age)
	fmt.Println(age)
	fmt.Println("快输入")
	fmt.Scanln(&sal)
	fmt.Println(sal)
	fmt.Println("快输入")
	fmt.Scanln(&isPass)
	fmt.Println(isPass)

	//method2
	fmt.Println("name,age,salary,isPass")
	fmt.Scanf("%s，%d,%f,%t",&name,&age,&sal,&isPass)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(sal)
	fmt.Println(isPass)


}
