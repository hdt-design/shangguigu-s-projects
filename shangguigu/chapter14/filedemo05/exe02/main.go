package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//打开一个存在的文件中，将原来的内容覆盖成新的内容10句 "你好，尚硅谷!"

	//创建一个新文件，写入内容 5句 "hello, Gardon"
	//1 .打开文件已经存在文件, /Users/liuxinpei/Desktop/abc.txt
	filepath := "/Users/liuxinpei/Desktop/abc.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	str := "你好，尚硅谷！\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		_, err := writer.WriteString(str)
		if err != nil {
			fmt.Println(err)
		}
	}
	if err := writer.Flush(); err != nil {
		fmt.Println(err)
	}
}
