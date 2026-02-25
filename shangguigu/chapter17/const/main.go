package main

import "fmt"

func main() {

	var num int
	num = 9 //ok
	//常量声明的时候必须赋值
	const tax int = 0
	//常量不能修改
	//tax=10
	fmt.Println(num, tax)
	//常量只能修饰bool，数值类型（int、float系列），string类型

	//fmt.Println(b)

	const (
		a = iota
		b
		c
		d
	)

	fmt.Println(a, b, c, d)
}
