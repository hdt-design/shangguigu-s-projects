package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point = Point{1, 2}
	a = point //ok
	//如何将a赋给point变量
	//b = a //error
	b := a.(Point)
	fmt.Println(b) //

	// //类型断言的其他案例
	// var x interface{}
	// var b2 float32 = 1.1
	// x = b2 //空接口，可接收任意类型
	// //x=>float32[使用类型断言]
	// y := x.(float32)
	// fmt.Printf("y的类型是%T 值是=%v", y, y)

	//类型断言（带检测）
	var x interface{}
	var b2 float32 = 1.1
	x = b2 //空接口，可接收任意类型
	//x=>float32[使用类型断言]

	//类型断言（带检测）
	if y, ok := x.(float32); ok {
		fmt.Println("convert success")
		fmt.Printf("y的类型是%T 值是=%v\n", y, y)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println("继续执行....")

}
