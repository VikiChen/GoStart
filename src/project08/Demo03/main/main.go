package main

import "fmt"

//map 初始化方式
func main() {

	studentMap :=make(map[string]map[string]string)
	studentMap["student1"]=make(map[string]string)
	studentMap["student1"]["mame"]="1"
	studentMap["student1"]["age"]="12"
	studentMap["student2"]=make(map[string]string)
	studentMap["student2"]["mame"]="12"
	studentMap["student2"]["age"]="123"

	for _,v :=range studentMap{
		fmt.Println(v)
	}

	delete(studentMap,"student1")

	for _,v :=range studentMap{
		fmt.Println(v)
	}
}
