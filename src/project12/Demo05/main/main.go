package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	file ,err:=os.OpenFile("d:/1.txt",os.O_WRONLY|os.O_TRUNC,0666)
	if err!=nil{
		fmt.Println("open file err=",err)
		return
	}
	defer file.Close()

	writer:=bufio.NewWriter(file)
	str:="你好\r\n"
	for i:=0;i<5;i++{
		writer.WriteString(str)
	}
	writer.Flush()


}