package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// CharCount 定义一个结构体，用于保存统计结果
type CharCount struct {
	ChCount    int //英文
	NumCount   int //数字
	SpaceCount int //空格
	OtherCount int //其他
}

func main() {

	//思路：打开一个文件，创一个reader
	//每读取一行，就去统计该行有多少个英文，数字，空格和其他字符
	//然后将结果保存到一个结构体
	fileName := "/Users/liuxinpei/Desktop/abc.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	//定义一个CharCount实例
	var count CharCount
	//创建一个reader
	reader := bufio.NewReader(file)

	//循环读取fileName的内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//遍历str
		for _, v := range str {

			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			default:
				count.OtherCount++
			}
		}
	}

	fmt.Printf("字符的个数为=%v 数字的个数为=%v 空格的个数为=%v 其它字符个数=%v",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
