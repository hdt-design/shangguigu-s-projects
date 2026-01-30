package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var intarr [3]int32
	fmt.Println(intarr)
	fmt.Printf("intarr的地址是：%p intarr[0]的地址是：%p intarr[1]的地址是：%p intarr[2]的地址是：%p\n",
		&intarr, &intarr[0], &intarr[1], &intarr[2])
	fmt.Println(unsafe.Sizeof(intarr[0]))
	fmt.Println(unsafe.Sizeof(intarr))

	var score [5]float64

	for i := 0; i < len(score); i++ {
		fmt.Printf("请输入第%d个元素的值：", i+1)
		fmt.Scanln(&score[i])
	}
	for i := 0; i < len(score); i++ {
		fmt.Printf("score[%d]=%v\n", i, score[i])
	}
	var numberarr01 = [3]int{1, 2, 3}
	fmt.Println("numberarr01:", numberarr01)

	var numberarr02 = [3]int{5, 6, 7}
	fmt.Println("numberarr02:", numberarr02)

	var numberarr03 = [...]int{8, 9, 10}
	fmt.Println("numberarr03:", numberarr03)

	var numberarr04 = [...]int{1: 100, 0: 300, 2: 500}
	fmt.Println("numberarr04:", numberarr04)
}
