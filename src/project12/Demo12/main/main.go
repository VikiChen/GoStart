package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user,"u","","用户名")
	flag.StringVar(&pwd,"p","","密码")
	flag.StringVar(&host,"h","","ip")
	flag.IntVar(&port,"port",0,"端口")
	flag.Parse()
	fmt.Print(user,"-",pwd,"-",host,"-",port)

}

