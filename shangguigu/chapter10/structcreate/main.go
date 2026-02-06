package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	//方式1

	//方式2
	p2 := Person{"mary", 20}
	// p2.Name = "Tom"
	// p2.Age = 18
	fmt.Println(p2)

	//方式3
	var p3 = new(Person)
	//var p3 *Person=new(Person)

	//因为p3是一个指针，因此标准的给字符赋值的方式
	//也可以写成p3.Name
	//原因：go的设计者为了程序员使用方便，底层会对p3.name进行处理
	//会给p3加上取值运算
	(*p3).Name = "smith"
	p3.Name = "jane"
	(*p3).Age = 30
	fmt.Println(p3)

	//方式4
	var person = &Person{}
	//var person *Person=&Person{}
	//因为person是一个指针，因此标准的给字符赋值的方式
	//（*person）.Name = "Tom"
	// person.Name = "Tom"(底层处理)
	(*person).Name = "scott"
	person.Name = "scott~~"
	(*person).Age = 88
	person.Age = 10
	fmt.Println(person)
}
