package main

import "fmt"

//map 切片
func main() {
	var mas []map[string]string
	mas=make([]map[string]string,2)
	mas[0]=make(map[string]string)
	mas[0]["name"]="ksjdks"
	mas[1]=make(map[string]string)
	mas[1]["name"]="ksjdks"
	newMOS :=map[string]string{"name":"sd",}
	mas=append(mas,newMOS)
	fmt.Println(mas[0])
}
