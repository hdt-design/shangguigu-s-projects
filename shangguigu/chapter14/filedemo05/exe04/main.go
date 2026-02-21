package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	//打开一个存在的文件，将原来的内容读出显示在终端，并且追加5句"hello,北京!"
	//1 .打开文件已经存在文件, /Users/liuxinpei/Desktop/abc.txt
	filepath := "/Users/liuxinpei/Desktop/abc.txt"
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("close file failed, err:%v\n", err)
		}
	}()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}

	str := "hello,北京！\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		_, err := writer.WriteString(str)
		if err != nil {
			fmt.Printf("write file failed, err:%v\n", err)
		}
	}
	if err := writer.Flush(); err != nil {
		fmt.Printf("flush file failed, err:%v\n", err)
	}
}
