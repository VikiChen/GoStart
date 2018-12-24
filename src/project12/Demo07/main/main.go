package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func main() {
	file ,err:=os.OpenFile("d:/1.txt",os.O_RDWR|os.O_APPEND,0666)
	if err!=nil{
		fmt.Println("open file err=",err)
		return
	}
	defer file.Close()
	newReader := bufio.NewReader(file)
	for {
		n, e := newReader.ReadString('\n')
		if e==io.EOF{
			break
		}
		fmt.Println(string(n))
	}
	writer:=bufio.NewWriter(file)
	str:="厉害2\r\n"
	for i:=0;i<5;i++{
		writer.WriteString(str)
	}
	writer.Flush()


}