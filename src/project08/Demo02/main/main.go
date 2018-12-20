package main

import "fmt"

//map 初始化方式
func main() {
	var a map[string]string=make(map[string]string,10)
	a["no1"]="ww"
	a["no2"]="2w"
	fmt.Println(a)

	cities :=make(map[string]string)
	cities["wdwd"]="bei"
}
