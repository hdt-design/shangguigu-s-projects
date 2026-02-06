package main

import (
	"fmt"
)

/*学生案例，编写一个student结构体，包含name,gender,age,id,score字段，
分别为string,string,int,int,float64类型。
在结构体中声明一个say方法，返回string类型，方法返回信息中包含所有字段值
在main方法中，创建student结构体实例（变量），并访问say方法，并将调用结果打印输出*/

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (student *Student) say() string {
	infoStr := fmt.Sprintf("学生信息：姓名：%s，性别：%s，年龄：%d，学号：%d，成绩：%.2f",
		student.name, student.gender, student.age, student.id, student.score)

	return infoStr
}

/*
1）编程创建一个Box结构体，在其中声明三个字段表示一个立方体的长，宽，高，从终端获取
2）声明一个方法获取立方体的体积
3）创建一个Box结构体变量，打印给定尺寸的立方体的体积
*/

type Box struct {
	len    float64
	width  float64
	height float64
}

func (box *Box) getVolume() float64 {
	return box.len * box.width * box.height
}

/*
	景区门票案例
	一个景区根据游人的年龄收取不同价格的门票，年龄大于18，收费20，其他情况门票免费
	编写visitor结构体，根据年龄段决定能够购买的门票价格并输出
*/

type Visitor struct {
	Name string
	Age  int
}

func (visitor *Visitor) showprice() {
	if visitor.Age > 18 {
		fmt.Printf("游客名字为%v 年龄为%v 收费20\n", visitor.Name, visitor.Age)
	} else {
		fmt.Printf("游客名字为%v 年龄为%v 免费\n", visitor.Name, visitor.Age)
	}
}

func main() {
	//测试
	//创建一个student实例
	var stu = Student{
		name:   "tom",
		gender: "male",
		age:    18,
		id:     1000,
		score:  99.98,
	}
	fmt.Println(stu.say())

	//测试Box结构体
	var box Box
	fmt.Println("请输入立方体的长、宽、高：")
	fmt.Scanln(&box.len, &box.width, &box.height)
	fmt.Printf("立方体的体积为：%.2f\n", box.getVolume())

	//测试visitor结构体
	var v Visitor
	for {
		fmt.Println("请输入你的名字")
		fmt.Scanln(&v.Name)
		if v.Name == "n" {
			fmt.Println("退出程序")
			break
		}
		fmt.Println("请输入你的年龄")
		fmt.Scanln(&v.Age)
		v.showprice()
	}
}
