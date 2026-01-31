package main

import "fmt"

func bubbleSort(arr *[5]int) {
	fmt.Println("排序前arr=", *arr)
	temp := 0
	for j := 0; j < len(arr)-1; j++ {
		for i := 0; i < len(arr)-j-1; i++ {
			if arr[i] > arr[i+1] {
				temp = arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
			}
		}
		fmt.Printf("第%d轮排序结果=%v\n", j+1, *arr)
	}
	fmt.Println("排序后arr=", *arr)
}

func main() {
	arr := [5]int{24, 69, 80, 57, 13}
	bubbleSort(&arr)

	fmt.Println("main函数中的arr=", arr)
}

// Output:
// 排序前arr=[24 69 80 57 13]
// 第1轮排序结果=[24 69 80 57 13]
// 第2轮排序结果=[24 57 69 80 13]
// 第3轮排序结果=[24 57 69 13 80]
// 第4轮排序结果=[24 57 13 69 80]
// 排序后arr=[13 24 57 69 80]
// main函数中的arr=[13 24 57 69 80]
