package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func CopyFile(dstFileName string,srcFileName string)(written int64,err error)  {
	srcfile,err:=os.Open(srcFileName)
	if err!=nil{
		fmt.Printf("open err=",err)
	}
	reader := bufio.NewReader(srcfile)

	file, e := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if e!=nil{
		fmt.Printf("open err=",e)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	return io.Copy(writer,reader)
}

func main() {
	file1Path :="d:/abd.txt"
	file2Path:="d:kkk.txt"
	written, err := CopyFile(file1Path, file2Path)
	if err !=nil{
		fmt.Println("copy err=",err)
	}
	fmt.Println(string(written))
}