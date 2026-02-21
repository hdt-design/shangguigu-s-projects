package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	//概念说明: file 的叫法
	//1. file 叫 file对象
	//2. file 叫 file指针
	//3. file 叫 file 文件句柄
	file, err := os.Open("/Users/liuxinpei/Desktop/text.txt")
	if err != nil {
		fmt.Println("create file err:", err)
		return
	}

	//当函数退出时，要及时的关闭file
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("close file err:", err)
		}
	}() //// 及时关闭 file，否则会导致文件描述符泄漏

	// 创建一个 *Reader  ，是带缓冲的
	/*
		const (
		defaultBuffSize = 4096 //默认的缓冲区为4096
		)
	*/
	reader := bufio.NewReader(file)
	//循环的读取文件的内容
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束
		if err == io.EOF {
			break
		}
		//输出内容
		fmt.Println(str)
	}

	fmt.Println("文件读取结束")
}
