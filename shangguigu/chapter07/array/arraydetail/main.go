package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// var arr01 [3]int
	// arr01[0] = 10
	// arr01[1] = 20
	// // arr01[2] = 20.1
	// // // 这里会报错，不能将float64赋值给int类型
	// arr01[2] = 30
	// // arr01[3]=40 // 这里会报错，数组越界
	// fmt.Println("arr01=", arr01)

	//数组创建后，如果没有赋值，会自动初始化为零值
	//1.数值（整数，浮点数） 0
	//2.布尔值 false
	//3.字符串 "" 空字符串

	var arr01 [3]float32
	var arr02 [3]string
	var arr03 [2]bool

	fmt.Printf("arr01=%v arr02=%v arr03=%v\n",
		arr01, arr02, arr03)

	// 输出：
	// arr01=[0 0 0] arr02=[  ] arr03=[false false]

	//请求出一个数组的和和平均值 for-range
	//1.声明一个数组 var intarr[5]
	//2.使用for-range遍历数组，求和和平均值
	//3.打印结果

	var intarr = [...]int{1, 2, 3, 4, 5}
	var sum = 0
	for _, v := range intarr {
		sum += v
	}
	avg := float32(sum) / float32(len(intarr))

	fmt.Printf("数组的和为：%d\n", sum)
	fmt.Printf("数组的平均值为：%.2f\n", avg)

	//随机生成五个数，并反转打印
	//1.随机生成五个数，rand.Intn(）函数
	//2.得到init数组
	//3.反转打印，交换次数是len/2

	var intarr3 [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		intarr3[i] = rand.Intn(100)
	}
	fmt.Println("随机生成的数组为：", intarr3)

	//反转打印
	temp := 0
	for i := 0; i < len(intarr3)/2; i++ {
		temp = intarr3[len(intarr3)-1-i]
		intarr3[len(intarr3)-1-i] = intarr3[i]
		intarr3[i] = temp
	}
	fmt.Println("反转后的数组为：", intarr3)
}
