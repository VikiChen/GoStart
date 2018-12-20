package main

import "fmt"

//map
func main() {
	var a map[string]string=make(map[string]string,10)
	a["no1"]="ww"
	a["no2"]="2w"
	fmt.Println(a)
}
