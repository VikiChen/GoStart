package main

import (
	"fmt"
)

func main() {
	key:=""
	loop :=true

	balance:=10000.0
	money:=0.0
	note:=""
	detail:="收支\t账户金额\t收支金额\t说明\n"
	flag:=false
	for  {
		fmt.Println("-------------------------家庭收支软件-----------------")
		fmt.Println("1.收支明细")
		fmt.Println("2.登记收入")
		fmt.Println("3.登记支出")
		fmt.Println("4.退出软件")
		fmt.Println("")
		fmt.Println("请选择（1-4)")

		fmt.Scanln(&key)
		switch key {
			case "1":
				fmt.Println("\n-----------当前收支记录-----------")
				if flag{
					fmt.Println(detail)
				}else {
					fmt.Println("没有收支快来一比")
				}
		case "2":
			fmt.Println("输入金额")
			fmt.Scanln(&money)
			fmt.Println("输入说明")
			fmt.Scanln(&note)
			balance+=money
			flag=true
			detail+=fmt.Sprintf("收入\t%v\t%v\t%v\n",balance,money,note)
		case "3":
			fmt.Println("输入金额")
			fmt.Scanln(&money)
			if money>balance{
				fmt.Println("余额不足\n")
				break
			}
			balance-=money
			fmt.Println("输入说明")
			fmt.Scanln(&note)
			flag=true
			detail+=fmt.Sprintf("支出\t%v\t%v\t%v\n",balance,money,note)
		case "4":
			fmt.Println("确定要退出吗y/n")
			choice:=""
			for{
				fmt.Scanln(&choice)
				if choice=="y"||choice=="n"{
					break
				}
				fmt.Println("输入有误重新输入")

			}
			if choice=="y"{
				loop=false
			}

		default:
			fmt.Println("请输入正确的选项")

		}
		if !loop{
			break
		}
	}
	fmt.Println("你退出了软件的使用")
}