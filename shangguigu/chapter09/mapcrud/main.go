package main

import (
	"fmt"
)

func main() {
	cities := make(map[string]string)
	cities["no1"] = "beijing"
	cities["no2"] = "shanghai"
	cities["no3"] = "guangzhou"
	fmt.Println(cities)

	//修改
	cities["no3"] = "tianjin"
	fmt.Println(cities)

	//删除
	delete(cities, "no2")
	fmt.Println(cities)
	//当指定key不存在时，删除操作不进行，也不报错main
	delete(cities, "no4")
	fmt.Println(cities)

	//查找
	val, ok := cities["no1"]
	if ok {
		fmt.Printf("有no1这个key，对应的值是%v\n", val)
	} else {
		fmt.Println("没有no1这个key")
	}

	//如果希望删除所有key
	//1.遍历所有key，逐个删除
	//2.make一个新的空间
	cities = make(map[string]string)
	fmt.Println(cities)
}
