package main
import (
	"fmt"
	"strings"
)

func Fscanf(r io.Reader, format string, a ...interface{})
(n int, err error)

file, err = os.Open("data.txt")
if err != nil {
    fmt.Println("无法打开文件:", err)
    return
}
defer file.Close()

var name string
var age int

// 从文件中读取数据
for {
    _, err := fmt.Fscanf(file, "%s %d\n", &name, &age)
    if err != nil {
        break // 读取完毕或出现错误
    }
    fmt.Printf("姓名: %s, 年龄: %d\n", name, age)
}