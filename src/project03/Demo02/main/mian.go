package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	j := 1
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	for {
		fmt.Println("ok")
		break
	}

	var str string = "hello world北京"
	str2 :=[]rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}

	for index, val := range str {
		fmt.Printf("index=%d,val=%c\n", index, val)
	}
}
