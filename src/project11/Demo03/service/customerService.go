package service

import "project11/Demo03/model"

type CustomerService struct {
	customers []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService  {

	customerService:=CustomerService{}
	customerService.customerNum=1
	customer :=model.NewCustomer(1,"张三","男",20,"222","ws@qq.com")
	customerService.customers=append(customerService.customers,customer)
	return &customerService
}

func (this *CustomerService) List()  []model.Customer{
	return this.customers
}

func (this *CustomerService) Add(customer model.Customer)  bool{
	this.customerNum+=1
	customer.Id=this.customerNum
	this.customers =append(this.customers, customer)
	return true
}

func (this *CustomerService) Delete(Id int)  bool{
	index :=this.FindById(Id)
	if index==-1{
		return false
	}
	this.customers=append(this.customers[:index],this.customers[index+1:]...)
	return true
}

func (this *CustomerService) FindById(Id int)  int{
	index :=-1
	for i,v:=range this.customers{
		if v.Id==Id{
			index=i
		}
	}
	return index
}

//
//func (this *CustomerService)   {
//
//}
