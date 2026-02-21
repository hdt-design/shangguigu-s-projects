package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//创建一个新文件，写入内容 5句 "hello, Gardon"
	//1 .打开文件 /Users/liuxinpei/Desktop/abc.txt
	filepath := "/Users/liuxinpei/Desktop/abc.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	//及时关闭file句柄
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	//准备写入5句 "hello, Gardon"
	str := "hello,Garden\r\n" // \r\n 表示换行
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		_, err := writer.WriteString(str)
		if err != nil {
			fmt.Println(err)
		}
	}
	//因为writer是带缓存，因此在调用WriterString方法时，其实
	//内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	//真正写入到文件中， 否则文件中会没有数据!!!
	if err = writer.Flush(); err != nil {
		fmt.Println(err)
	}
}
