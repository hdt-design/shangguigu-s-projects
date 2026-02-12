package main

import "fmt"

type Binterface interface {
	test01()
}

type Cinterface interface {
	test02()
}

type Ainterface interface {
	Binterface
	Cinterface
	test03()
}

// 1.如果需要实现AInterface,就需要将BInterface CInterface的方法都实现

type Stu struct {
}

func (stu Stu) test01() {

}

func (stu Stu) test02() {

}

func (stu Stu) test03() {

}

//2.可以将任何数据类型赋给空接口

type T interface {
}

func main() {
	var stu Stu
	var a Ainterface = stu
	a.test01()

	var t T = stu
	fmt.Println(t)
	var t2 interface{} = stu
	var num1 = 8.8
	t2 = num1
	t = num1
	fmt.Println(t2, t)
}
