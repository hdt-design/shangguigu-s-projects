package main

import (
	"fmt"
)

func main() {

	//演示切片的基本使用
	var intarr = [...]int{1, 2, 3, 4, 5}
	//声明一个切片
	//slice:=intarr[1:4]
	//1.slice就是切片名
	//2.intarr【1:4】表示slice引用到intarr这个数组
	//3.引用intarr数组的从下标1开始，一直到下标4（不包含下标4）的元素
	slice := intarr[1:4]
	fmt.Println("intarr=", intarr)
	fmt.Println("slice=", slice)
	fmt.Println("slice长度为：", len(slice))
	fmt.Println("slice容量为：", cap(slice))
	//切片容量可以动态变化

	fmt.Printf("intarr的地址=%p\n", &intarr[1])
}
