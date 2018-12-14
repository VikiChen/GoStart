package main

import "fmt"

func main() {
	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte ='+'
	result :=cal(n1,n2,operator)
	fmt.Println(result)
}
func cal(f float64, f2 float64, b byte) float64 {
	var res float64
	switch b {
	case '+':
		res=f+f2
	case '-':
		res=f-f2
	}
	return res
}