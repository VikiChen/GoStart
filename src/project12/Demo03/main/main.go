package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func main() {
	file ,err:=os.Open("d:/1.txt")
	if err!=nil{
		fmt.Println("open file err=",err)
	}
//	fmt.Println("file=",file)
	defer file.Close()
	//err2:=file.Close()
	//if err!=nil{
	//	fmt.Println("err=",err2)
	//}
	//size 4096
	reader:=bufio.NewReader(file)
	for{
		readString, err := reader.ReadString('\n')
		if err==io.EOF{
			break
		}
		fmt.Println(readString)
	}
	fmt.Println("读取结束")
}