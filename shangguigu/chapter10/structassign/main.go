package main

import "fmt"

type Stu struct {
	Name string
	Age  int
}

func main() {

	//创建结构体时直接指定字段值
	var stu1 = Stu{"小明", 19}
	stu2 := Stu{"小明～", 20}

	//在创建结构体变量时，把字段名和字段值写在一起,这种写法不依赖于字段的定义顺序
	var stu3 = Stu{
		Name: "jack",
		Age:  21,
	}
	stu4 := Stu{
		Age:  30,
		Name: "mary",
	}

	fmt.Println(stu1, stu2, stu3, stu4)

	//方式2,返回英国指针类型
	var stu5 = &Stu{"小红", 22}
	stu6 := &Stu{"小红～", 23}
	//在创建结构体指针变量时，把字段名和字段值写在一起,这种写法不依赖于字段的定义顺序
	var stu7 = &Stu{
		Name: "lucy",
		Age:  24,
	}
	stu8 := &Stu{
		Age:  31,
		Name: "tom",
	}

	fmt.Println(stu5, stu6, stu7, stu8)
}
