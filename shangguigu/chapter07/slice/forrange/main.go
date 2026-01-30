package main

import (
	"fmt"
)

func main() {
	//演示切片的遍历方式1：for range
	var arr = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	for j := 0; j < len(slice); j++ {
		fmt.Printf("slice[%v]=%v\n", j, slice[j])
	}

	//输出：
	//slice[0]=20
	// slice[1]=30
	// slice[2]=40

	fmt.Println("--------for range-------")
	for i, v := range slice {
		fmt.Printf("i=%v,v=%v\n", i, v)
	}

	//输出：
	//i=0,v=20
	//i=1,v=30
	//i=2,v=40

	fmt.Println()
	slice2 := slice[1:2]
	fmt.Println("slice2=", slice2)
	fmt.Println("slice=", slice)
	fmt.Println("arr", arr)

	//输出：
	//slice2=[30]
	//slice=[20 30 40]
	//arr [10 20 30 40 50]

	fmt.Println()
	var slice3 = []int{100, 200, 300}
	slice3 = append(slice3, 400, 500, 600) //追加元素

	slice3 = append(slice3, slice3...) //切片追加切片

	fmt.Println("slice3=", slice3)

	//输出：
	//slice3=[100 200 300 400 500 600 100 200 300 400 500 600]

	//切片的拷贝
	//copy
	fmt.Println()
	var slice4 = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	copy(slice5, slice4)
	fmt.Println("slice5=", slice5)
	fmt.Println("slice4=", slice4)

	//输出：
	//slice5=[1 2 3 4 5 0 0 0 0 0]
	//slice4=[1 2 3 4 5]
}
