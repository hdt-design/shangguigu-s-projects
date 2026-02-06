package main

import "fmt"

// 定义一个Cat结构体

type Cat struct {
	Name  string
	Age   int
	Color string
}

func main() {
	//张老太太养了2只猫，一只叫“小白”一只叫“小黑”
	//输入小猫名字时输出猫的名字和颜色main

	// //1.使用变量处理
	// var cat1name string = "小白"
	// var cat1color string = "白色"
	// var cat1age int = 3
	// var cat2name string = "小黑"
	// var cat2color string = "黑色"
	// var cat2age int = 2

	// //2.使用数组解决
	// var cainames [2]string = [...]string{"小白", "小黑"}
	// var catages [2]int = [...]int{3, 2}
	// var catcolors [2]string = [...]string{"白色", "黑色"}

	//3.使用结构体解决

	//创建一个cat变量
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	fmt.Println(cat1)
	fmt.Println("猫的信息如下")
	fmt.Println("name=", cat1.Name)
	fmt.Println("age=", cat1.Age)
	fmt.Println("color=", cat1.Color)
}
