package main

import (
	"fmt"
)

func main() {
	//二维数组
	/*
		0 0 0 0 0 0
		0 0 1 0 0 0
		0 2 0 3 0 0
		0 0 0 0 0 0
	*/
	//声明一个二维数组
	var arr [4][6]int
	//初始化二维数组
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	//遍历二维数组
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}

}
