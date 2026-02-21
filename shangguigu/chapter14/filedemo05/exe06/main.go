package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// CopyFile 自己编写一个函数，接收两个文件路径 srcFileName dsFileNAme
func CopyFile(dsFileName string, srcFileName string) (written int64, err error) {

	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := srcFile.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	reader := bufio.NewReader(srcFile)
	dstFile, err := os.OpenFile(dsFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := bufio.NewWriter(dstFile)
	defer func() {
		if err := dstFile.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	written, err = io.Copy(writer, reader)
	if err = writer.Flush(); err != nil {
		fmt.Println(err)
	}
	return written, err
}
func main() {

	//将/Users/liuxinpei/Downloads/IMG_1633.JPG文件拷贝到/Users/liuxinpei/Desktop/abc.JPG

	//调用CopyFile函数
	srcFile := "/Users/liuxinpei/Downloads/IMG_1633.JPG"
	dstFile := "/Users/liuxinpei/Desktop/abc.JPG"
	_, err := CopyFile(dstFile, srcFile)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("拷贝完成")
	}
}
