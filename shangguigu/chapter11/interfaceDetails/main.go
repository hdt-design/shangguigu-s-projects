package main

import (
	"fmt"
)

// 1.接口本身不能创造一个实例（say），但可以指向指向一个实现了该接口的自定义类型的变量

type Ainterface interface {
	Say()
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}

// 2.指向自定义类型的变量不一定要是结构体，只要实现了该接口的类型都可以
type integer int

func (i integer) Say() {
	fmt.Println("integer Say i =", i)
}

// 3.同一个自定义类型的变量可以实现多个接口

type Binterface interface {
	Hello()
}

type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("Hello Monster")
}

func (m Monster) Say() {
	fmt.Println("Monster Say")
}

func main() {
	var stu Stu
	var a Ainterface = stu
	a.Say()

	var i integer = 10
	var b Ainterface = i
	b.Say()

	//Monster实现了Ainterface和Binterface
	var monster Monster
	var a2 Ainterface = monster
	var b2 Binterface = monster
	a2.Say()
	b2.Hello()
}
