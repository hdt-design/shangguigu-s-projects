package main

import (
	"fmt"
)

//二分查找

func binaryfind(arr *[6]int, leftindex int, rightindex int, findval int) {
	if leftindex > rightindex {
		fmt.Println("没有找到")
		return
	}

	middle := (leftindex + rightindex) / 2
	if (*arr)[middle] > findval {
		binaryfind(arr, leftindex, middle-1, findval)
	} else if (*arr)[middle] < findval {
		binaryfind(arr, middle+1, rightindex, findval)
	} else {
		fmt.Printf("找到了,下标为%v\n", middle)
	}
}
func main() {
	arr := [6]int{1, 8, 10, 89, 1000, 1234}
	binaryfind(&arr, 0, len(arr)-1, 1000)
}
