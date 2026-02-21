package main

import (
	"fmt"
	"os"
)

func main() {

	//将/Users/liuxinpei/Desktop/abc.txt文件内容导入到/Users/liuxinpei/Desktop/kkk.txt
	//1.首先将/Users/liuxinpei/Desktop/abc.txt内容读取到内存
	//2.将读取到的内容写入/Users/liuxinpei/Desktop/kkk.txt
	filepath := "/Users/liuxinpei/Desktop/abc.txt"
	file2path := "/Users/liuxinpei/Desktop/kkk.txt"
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.WriteFile(file2path, data, 0666)
	if err != nil {
		fmt.Println(err)
	}
}
