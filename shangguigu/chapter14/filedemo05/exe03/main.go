package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//打开一个存在的文件，在原来的内容追加内容 'ABC! ENGLISH!'
	//1 .打开文件已经存在文件, /Users/liuxinpei/Desktop/abc.txt
	filepath := "/Users/liuxinpei/Desktop/abc.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	str := "ABC,ENGLISH\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		_, err := writer.WriteString(str)
		if err != nil {
			fmt.Println(err)
		}
	}
	if err := writer.Flush(); err != nil {
		fmt.Println(err)
	}
}
