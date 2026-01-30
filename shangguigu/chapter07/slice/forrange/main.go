package main

import (
	"fmt"
)

func main() {
	//演示切片的遍历方式1：for range
	var arr = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v\n", i, slice[i])
	}

	fmt.Println("--------for range-------")
	for i, v := range slice {
		fmt.Printf("i=%v,v=%v\n", i, v)
	}

	fmt.Println()
	slice2 := slice[1:2]
	fmt.Println("slice2=", slice2)
	fmt.Println("slice=", slice)
	fmt.Println("arr", arr)

	fmt.Println()
	var slice3 = []int{100, 200, 300}
	slice3 = append(slice3, 400, 500, 600) //追加元素

	slice3 = append(slice3, slice3...) //切片追加切片

	fmt.Println("slice3=", slice3)
}
