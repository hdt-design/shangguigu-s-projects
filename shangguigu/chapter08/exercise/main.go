package main

import (
	"fmt"
)

func main() {
	/*定义二维数组，用于保存三个班，每个班五名同学成绩
	并求出班级平均分，所有班级平均分
	*/

	//定义二维数组
	var score [3][5]float64
	//2.循环输入成绩
	for i := 0; i < len(score); i++ {
		for j := 0; j < len(score[i]); j++ {
			fmt.Print("请输入第", i+1, "个班第", j+1, "名同学的成绩:")
			fmt.Scan(&score[i][j])
		}
	}

	//fmt.Println(score)

	//3.遍历输出成绩后的二维数组，统计平均分
	totalsum := 0.0 //总分
	for i := 0; i < len(score); i++ {
		sum := 0.0
		for j := 0; j < len(score[i]); j++ {
			sum += score[i][j]
		}
		totalsum += sum
		fmt.Printf("第%d个班的总分为:", i+1)
		fmt.Printf("第%d个班的平均分为%.2f\n", i+1,
			sum/float64(len(score[i])))
	}
	//4.输出所有班级平均分
	fmt.Printf("所有班级的平均分为%.2f\n",
		totalsum/float64(len(score)*len(score[0])))
}
