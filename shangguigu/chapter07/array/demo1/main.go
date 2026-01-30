package main

import (
	"fmt"
)

func main() {
	//使用数组的方式解决问题
	//定义一个数组
	var hens [6]float64
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0
	//遍历数组
	totalweight := 0.0
	for i := 0; i < len(hens); i++ {
		totalweight += hens[i]
	}
	avgweight := totalweight / float64(len(hens))
	fmt.Printf("totalweight=%.2f, avgweight=%.2f\n", totalweight, avgweight)
}
