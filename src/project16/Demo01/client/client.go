package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func main() {
	fmt.Println("开始连接")
	conn, e := net.Dial("tcp", "127.0.0.1:8988")
	if e!=nil{
		fmt.Println("连接异常",e)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		s, i := reader.ReadString('\n')
		strings.Trim(s," \r\n")
		if s=="exit"{
			fmt.Println("客户端退出")
			break
		}
		if i != nil {
			fmt.Println("读取失败")
		}
		_, err := conn.Write([]byte(s+"\n"))
		if err != nil {
			fmt.Println("connect write err", err)
		}
	}

}
