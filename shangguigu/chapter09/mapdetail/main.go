package main

import (
	"fmt"
)

func modify(map1 map[int]int) {
	map1[10] = 900
}

// 定义一个学生结构体
type stu struct {
	name    string
	age     int
	address string
}

func main() {
	//map是引用类型，就难受引用类型传递的机制，在一个函数接收map
	//修改后，会直接修改原来的map

	map1 := make(map[int]int)
	map1[1] = 90
	map1[2] = 88
	map1[10] = 1
	map1[20] = 2
	modify(map1)
	fmt.Println(map1)

	//map的value也经常使用struct类型
	//更适合管理复杂的数据（比前面value是一个map更好）
	//比如value为student结构体
	//1. map的key为学生学号
	//2. map的value为学生结构体，包含名字，年龄等

	students := make(map[string]stu, 10)
	//创建2个学生
	stu1 := stu{"张三", 20, "北京"}
	stu2 := stu{"李四", 28, "上海"}
	students["no1"] = stu1
	students["no2"] = stu2

	fmt.Println(students)

	//遍历学生信息
	for k, v := range students {
		fmt.Printf("学生等编号是%v\n", k)
		fmt.Printf("学生的名字是%v\n", v.name)
		fmt.Printf("学生的年龄是%v\n", v.age)
		fmt.Printf("学生的地址是%v\n", v.address)
		fmt.Println()
	}
}
