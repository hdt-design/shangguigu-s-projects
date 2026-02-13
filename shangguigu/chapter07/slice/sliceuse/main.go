package main

import (
	"fmt"
)

func main() {
	//演示切片的基本使用
	var slice = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 20
	//对于切片必须slice
	fmt.Println(slice)
	fmt.Println("slice的长度=", len(slice))
	fmt.Println("slice的容量=", cap(slice))

	fmt.Println()
	var strslice = []string{"hello", "world"}
	fmt.Println(strslice)
	fmt.Println("strslice的长度=", len(strslice))
	fmt.Println("strslice的容量=", cap(strslice))
}
