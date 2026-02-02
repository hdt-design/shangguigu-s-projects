package main

import (
	"fmt"
)

func main() {
	//使用for-range循环遍历map
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "广州"

	for k, v := range cities {
		fmt.Printf("k=%v v=%v\n", k, v)
	}

	//Output
	//k=no1 v=北京
	//k=no2 v=上海
	//k=no3 v=广州

	//使用for-range遍历一个结构复杂的map
	studentmap := make(map[string]map[string]string)
	studentmap["student1"] = make(map[string]string, 3)
	studentmap["student1"]["name"] = "张三"
	studentmap["student1"]["age"] = "20"
	studentmap["student1"]["gender"] = "男"

	studentmap["student2"] = make(map[string]string, 3)
	studentmap["student2"]["name"] = "李四"
	studentmap["student2"]["age"] = "21"
	studentmap["student2"]["gender"] = "女"

	for k1, v1 := range studentmap {
		fmt.Println(k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t%v:%v\n", k2, v2)
		}
		fmt.Println()
	}

	//Output
	//student1
	//	name:张三
	//	age:20
	//	gender:男
	//
	//student2
	//	name:李四
	//	age:21
	//	gender:女

}
