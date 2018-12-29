package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn)  {
	defer conn.Close()
	for{
		buf :=make([]byte,1024)
		fmt.Println("服务器在等待客户端发送信息 "+conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err!=nil{
			fmt.Println("客户端已经退出")
			return
		}
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听")
	listener, e := net.Listen("tcp", "0.0.0.0:8988")
	if e!=nil{
		fmt.Println("失败",e)
		return
	}
	defer listener.Close()

	for{
		fmt.Println("等待客户端来连接")
		conn, i := listener.Accept()
		if i!=nil{
			fmt.Println("获取连接失败",i)
		}else {
			fmt.Println(conn.RemoteAddr().String())
		}
		go process(conn)

	}

}
