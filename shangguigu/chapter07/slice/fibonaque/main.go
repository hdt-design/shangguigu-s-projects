package main

import (
	"fmt"
)

func fbn(n int) []uint64 {
	fnbslice := make([]uint64, n)
	fnbslice[0] = 1
	fnbslice[1] = 1
	for i := 2; i < n; i++ {
		fnbslice[i] = fnbslice[i-1] + fnbslice[i-2]
	}
	return fnbslice
}

func main() {
	/*
		1) can accept a n int
		2) put the Fibonacci sequence into a slice
		2）将斐波那契数列放入切片
		声明函数fbn（n int）([]uint64)
		进行for循环，放入切片
		返回切片
	*/

	//测试
	fmt.Println(fbn(10))
}
