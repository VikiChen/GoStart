package main

import "fmt"

type Account struct {
	AccountName string
	Pwd string
	Balance float64
}

func (account Account) Deposite(money float64 ,pwd string)  {
	if pwd !=account.Pwd{
		fmt.Println("密码不正确！")
		return
	}

	if money <=0{
		fmt.Println("金额不正确")
		return
}
	account.Balance +=money
	fmt.Println("存款成功")

}

//抽象
func main() {

}
