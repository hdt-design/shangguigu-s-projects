package main

import (
	"fmt"
)

func main() {
	//顺序查找
	//思路
	//1. 定义一个数组
	//2.控制台输入

	names := [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	var heroname = " "
	fmt.Println("请输入要查找的英雄名字:")
	fmt.Scanln(&heroname)

	//顺序查找：第一种方式
	for i := 0; i < len(names); i++ {
		if names[i] == heroname {
			fmt.Printf("找到了%v,在第%d个位置\n", heroname, i)
			break
		}
		if i == len(names)-1 {
			fmt.Printf("很遗憾，没有找到%v\n", heroname)
		}
	}

	//顺序查找：第二种方式（推荐）
	index := -1
	for i := 0; i < len(names); i++ {
		if names[i] == heroname {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Printf("很遗憾，没有找到%v\n", heroname)
	} else {
		fmt.Printf("找到了%v,在第%d个位置\n", heroname, index)
	}
}
