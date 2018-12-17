package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "hello！被"
	fmt.Println(len(str))

	str2 := []rune(str)
	for i := 0; i < len(str2); i++ { //字符切割会有问题
		fmt.Printf("%c\n", str2[i])
	}

	strx, err := strconv.Atoi("222")
	if err != nil {
		fmt.Println(strx)
	} else {
		fmt.Println("fail")
	}

	stry := strconv.Itoa(12)
	fmt.Println(stry)

	var b = []byte("hello go")
	fmt.Println(b)
}
