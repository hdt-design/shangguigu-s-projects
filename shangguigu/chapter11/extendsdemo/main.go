package main

import (
	"fmt"
)

//编写学生考试系统

type Student struct {
	Name  string
	Age   int
	Score int
}

//将Pupil和Graduate共有的方法也绑定到*Student

func (stu *Student) ShowInfo() {
	fmt.Printf("学生姓名：%s，年龄：%d，分数：%d\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	//业务判断
	stu.Score = score
}

//给*Student增加一个方法，pupil和Graduate都可以使用

func (stu *Student) GetSum(n1 int, n2 int) int {
	return n1 + n2
}

//小学生

type Pupil struct {
	Student
}

//显示他的成绩

//pupil特有的方法，保留

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中...")
}

//大学生，研究生

//大学生

type Graduate struct {
	Student
}

//显示他的成绩

//graduate特有的方法，保留

func (g *Graduate) testing() {
	fmt.Println("大学生正在考试中...")
}

func main() {
	//当我们对结构体嵌入匿名结构体使用方法会发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "tom~"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.SetScore(70)
	pupil.ShowInfo()
	fmt.Println(pupil.GetSum(1, 2))

	graduate := &Graduate{}
	graduate.Student.Name = "jerry~"
	graduate.Student.Age = 25
	graduate.testing()
	graduate.SetScore(80)
	graduate.ShowInfo()
	fmt.Println(graduate.GetSum(10, 20))
}
