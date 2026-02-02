package main

import (
	"fmt"
)

func main() {
	//第一种方式
	var a map[string]string
	a = make(map[string]string)
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no3"] = "武松"
	a["no4"] = "吴用"
	fmt.Println(a)

	//第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "广州"
	fmt.Println(cities)

	//第三种方式
	heroes := map[string]string{
		"hero1": "索隆",
		"hero2": "亚索",
		"hero3": "艾斯",
	}

	fmt.Println(heroes)

	//Output

	//案例
	//演示一个key-value存储的应用场景：存储用户信息
	studentmap := make(map[string]map[string]string)
	studentmap["student1"] = make(map[string]string, 3)
	studentmap["student1"]["name"] = "张三"
	studentmap["student1"]["age"] = "20"
	studentmap["student1"]["gender"] = "男"

	studentmap["student2"] = make(map[string]string, 3)
	studentmap["student2"]["name"] = "李四"
	studentmap["student2"]["age"] = "21"
	studentmap["student2"]["gender"] = "女"

	fmt.Println(studentmap)
}
