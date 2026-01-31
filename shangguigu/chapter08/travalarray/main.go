package main

import "fmt"

func main() {
	//演示二维数组的遍历
	var arr3 = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	//for循环
	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[i]); j++ {
			fmt.Printf("%v\t", arr3[i][j])
		}
		fmt.Println()
	}

	// for-range循环
	for i, v := range arr3 {
		for j, v2 := range v {
			fmt.Printf("arr3[%d][%d]=%d\n", i, j, v2)
		}
		fmt.Println()
	}
}

// 输出结果：
// 1       2       3
// 4       5       6
// arr3[0][0]=1
// arr3[0][1]=2
// arr3[0][2]=3
// arr3[1][0]=4
// arr3[1][1]=5
// arr3[1][2]=6
