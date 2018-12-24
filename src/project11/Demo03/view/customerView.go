package main

import (
	"fmt"
	"project11/Demo03/service"
	"project11/Demo03/model"
)

type customerView struct {
	key string
	loop bool
	customerService *service.CustomerService

}

func (this *customerView) mainMenu()  {
	for {

		fmt.Println("-------------------------客户信息-----------------")
		fmt.Println("1.添加客户")
		fmt.Println("2.修改客户")
		fmt.Println("3.客户删除")
		fmt.Println("4.客户列表")
		fmt.Println("5.退出")
		fmt.Println("请选择（1-5)")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			this.delete()
		case "3":
			//this.pay()
		case "4":
			this.list()
		case "5":
			this.loop=false
		default:
			fmt.Println("请输入正确的选项")
		}
		if !this.loop{
			break
		}
	}
	fmt.Println("成功退出")
}

func (this *customerView) list()  {
	customers:=this.customerService.List()
	fmt.Println("-----------------客户列表-------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱\t\n")
	for _,v :=range customers{
		fmt.Println(v.GetInfo())
	}

	fmt.Println("------------------客户列表结束------------")
}

func (this *customerView) add()  {
	fmt.Println("---------------------添加客户--------------")
	fmt.Println("输入姓名")
	 name :=""
	 fmt.Scanln(&name)
	fmt.Println("性别")
	gender :=""
	fmt.Scanln(&gender)
	fmt.Println("年龄")
	age :=0
	fmt.Scanln(&age)
	fmt.Println("电话")
	phone :=""
	fmt.Scanln(&phone)
	fmt.Println("邮箱")
	email :=""
	fmt.Scanln(&email)
	customer :=model.NewCustomer2(name,gender,age,phone,email)
	if this.customerService.Add(customer){
		fmt.Println("-------------添加完成----------------")
	}

}

func (this *customerView) delete()  {
	fmt.Println("------------------删除客户-----------------")
	fmt.Println("请选择删除客户的编号 （-1 退出）")
	id:=-1
	fmt.Scanln(&id)
	if id==-1{
		return  //放弃删除操作
	}
	fmt.Println("是否删除y/n")
	choice :=""
	fmt.Scanln(&choice)
	if choice =="y"||choice=="Y"{
		if this.customerService.Delete(id){
			fmt.Println("删除完成")
		}else {
			fmt.Println("删除失败id 不存在")
		}
	}



}

func main()  {
	customerView :=customerView{
		key:"",
		loop:true,
	}
	customerView.customerService=service.NewCustomerService()
	customerView.mainMenu()
}
