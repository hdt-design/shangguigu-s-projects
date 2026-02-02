package main

import (
	"fmt"
)

func main() {
	//演示map切片main
	/*要求：使用一个map来记录monster的信息name和age，也就是说
	一个monster对应一个map，并且monster的个数可以动态增加=>map切片*/
	//1.声明一个切片
	monsters := make([]map[string]string, 2) //两个monster
	//2.增加第一个monster的信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "1000"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string)
		monsters[1]["name"] = "玉兔精"
		monsters[1]["age"] = "500"
	}

	//会越界
	// if monsters[2] == nil {
	// 	monsters[2] = make(map[string]string)
	// 	monsters[2]["name"] = "白骨精"
	// 	monsters[2]["age"] = "300"
	// }

	//使用append增加
	//1.先定义monster
	newmonster := map[string]string{
		"name": "白骨精",
		"age":  "300",
	}
	monsters = append(monsters, newmonster)

	fmt.Println(monsters)
}
