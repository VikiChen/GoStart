package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

type CharCount struct {
	CharCount int
	NumCount int
	SpaceCount int
	OtherCount int
}


func main() {
	srcfile,err:=os.Open("d:/1.txt")
	if err!=nil{
		fmt.Printf("open err=",err)
		return
	}
	defer srcfile.Close()
	var count CharCount
	reader := bufio.NewReader(srcfile)
	for{
		s, e := reader.ReadString('\n')
		if e == io.EOF {
			break
		}
		for _,v :=range s{
			switch {
			case v>='a' && v<='z':
				fallthrough
			case v>='A' && v<='Z':
				count.CharCount++
			case v>=' ' && v<='\t':
				count.SpaceCount++
			case v>=' ' && v<='\t':
				count.NumCount++
			default:
				count.OtherCount++
		}
		}
		}
		fmt.Println(count)
	}

