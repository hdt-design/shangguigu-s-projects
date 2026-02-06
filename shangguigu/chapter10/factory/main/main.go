package main

import (
	"fmt"
	"main/shangguigu/chapter10/factory/model"
)

func main() {
	//创建实例
	// var stu = model.Student{
	// 	Name:  "tom",
	// 	Score: 78.9,
	// }

	//当student结构体首字母小写，通过工厂模式解决
	var stu = model.Newstudent("tom~", 88.8)
	//输出实例信息
	fmt.Println(*stu)
	fmt.Println("name=", stu.Name)
}
