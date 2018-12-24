package model

import "fmt"

type FamilyAccount struct {
	key string
	loop bool
	balance float64
	money float64
	note string
	detail string
	flag bool
}

func NewFamilyAccount() *FamilyAccount  {
	return &FamilyAccount{
		key:"",
		loop :true,
		balance:10000.0,
		money:0.0,
		note:"",
		detail:"收支\t账户金额\t收支金额\t说明\n",
		flag:false,
	}
}

func (this *FamilyAccount)Mainmenu() {
	for  {
		fmt.Println("-------------------------家庭收支软件-----------------")
		fmt.Println("1.收支明细")
		fmt.Println("2.登记收入")
		fmt.Println("3.登记支出")
		fmt.Println("4.退出软件")
		fmt.Println("")
		fmt.Println("请选择（1-4)")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDeatils()
		case "2":
			this.inCome()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项")
		}
		if !this.loop{
			break
		}
	}

}

func (this *FamilyAccount) showDeatils() {
	fmt.Println("\n-----------当前收支记录-----------")
	if this.flag{
		fmt.Println(this.detail)
	}else {
		fmt.Println("没有收支快来一比")
	}
}
func (this *FamilyAccount)  inCome() {
	fmt.Println("输入金额")
	fmt.Scanln(&this.money)
	fmt.Println("输入说明")
	fmt.Scanln(&this.note)
	this.balance+=this.money
	this.flag=true
	this.detail+=fmt.Sprintf("收入\t%v\t%v\t%v\n",this.balance,this.money,this.note)

}

func (this *FamilyAccount) pay()  {
	fmt.Println("输入金额")
	fmt.Scanln(&this.money)
	if this.money>this.balance{
		fmt.Println("余额不足\n")
	}
	this.balance-=this.money
	fmt.Println("输入说明")
	fmt.Scanln(&this.note)
	this.flag=true
	this.detail+=fmt.Sprintf("支出\t%v\t%v\t%v\n",this.balance,this.money,this.note)
}

func (this *FamilyAccount)exit()  {
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
		this.loop=false
	}
}