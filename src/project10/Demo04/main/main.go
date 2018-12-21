package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {

}

type Camera struct {

}

func (p Phone) Start()  {
	fmt.Println("手机开始工作")
}

func (p Phone) Stop()  {
	fmt.Println("手机停止工作")
}

func (p Camera) Start()  {
	fmt.Println("camera开始工作")
}

func (p Camera) Stop()  {
	fmt.Println("camera停止工作")
}

type Computer struct {

}

func (c Computer) Working(usb Usb)  {
	usb.Start()
	usb.Stop()
}
//接口
func main() {
	c:=Computer{}
	phone := Phone{}
	camera := Camera{}
	c.Working(phone)
	c.Working(camera)
}
