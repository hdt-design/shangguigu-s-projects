package main

import (
	"fmt"
)

//声明/定义一个接口

type Usb interface {
	//声明两个未实现的方法
	Start()
	Stop()
}

type Usb2 interface {
	Start()
	Stop()
	Test()
}

type Phone struct {
}

//让Phone实现USB接口的方法

func (p Phone) Start() {
	fmt.Println("手机开始工作")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

type Camera struct {
}

//让camera实现USB接口的方法

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

//计算机

type Computer struct {
}

//编写一个方法working，接收一个usb接口类型变量
//只要实现了usb接口（所谓实现usb接口，就是指实现了usb接口声明的所有方法）

func (c Computer) Working(usb Usb) {

	//通过usb接口变量来调用
	usb.Start()
	usb.Stop()
}

func main() {
	//测试
	//创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//关键
	computer.Working(phone)
	computer.Working(camera)
}
