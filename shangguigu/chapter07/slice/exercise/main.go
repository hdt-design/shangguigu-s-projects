package main

import (
	"fmt"
)

func main() {
	// //演示切片的基本使用
	// var intarr = [...]int{1, 2, 3, 4, 5}
	// //声明一个切片
	// //slice:=intarr[1:4]
	// //1.slice就是切片名
	// //2.intarr【1:4】表示slice引用到intarr这个数组
	// //3.引用intarr数组的从下标1开始，一直到下标4（不包含下标4）的元素
	// slice := intarr[1:4]
	// fmt.Println("intarr=", intarr)
	// fmt.Println("slice=", slice)
	// fmt.Println("slice长度为：", len(slice))
	// fmt.Println("slice容量为：", cap(slice))
	// //切片容量可以动态变化

	var slice []int
	var arr = [...]int{1, 2, 3, 4, 5}
	slice = arr[:]
	slice[0] = 10
	fmt.Println("slice=", slice)
	slice2 := arr[0:]
	fmt.Println("slice2", slice2)
	fmt.Println("arr", arr)

	//输出：
	//slice2 [10 2 3 4 5]
	//slice [10 2 3 4 5]
	// //arr [10 2 3 4 5]

	// fmt.Println()
	// //string底层是一个byte数组，因此string可以切片
	// str := "hello@world"
	// //使用切片获取world
	// slice3 := str[6:]
	// fmt.Println("slice3=", slice3)

	// //输出：
	// //slice3= world

	// //如果需要修改字符串，可以先将string-->[]byte切片
	// //或者[]rune切片，修改后再转换回string
	// arr1 := []byte(str)
	// arr1[0] = 'z'
	// str2 := string(arr1)
	// fmt.Println("str2=", str2)
}
