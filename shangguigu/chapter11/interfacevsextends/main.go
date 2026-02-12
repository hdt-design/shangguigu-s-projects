package main

import (
	"fmt"
)

//Monkey结构体

type Monkey struct {
	Name string
}

//声明接口

type BiedAble interface {
	Flying()
}

type fishAble interface {
	Swimming()
}

func (m *Monkey) climbing() {
	fmt.Println(m.Name, "生来会爬树")
}

//让LittleMonkey实现BiedAble接口

func (m *LittleMonkey) Flying() {
	fmt.Println(m.Name, "通过学习，会飞翔...")
}

//让LittleMonkey实现fishAble接口

func (m *LittleMonkey) Swimming() {
	fmt.Println(m.Name, "通过学习，会游泳...")
}

//LittleMonkey结构体

type LittleMonkey struct {
	Monkey //继承
}

func main() {

	//创建一个LittleMonkey实例
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}
