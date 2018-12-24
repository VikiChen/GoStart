package main

import (
	"project11/Demo02/model"
	"fmt"
)

func main() {
	familyAccount:=model.NewFamilyAccount()
	for {
		familyAccount.Mainmenu()
	}
	fmt.Println("你退出了软件的使用")
}