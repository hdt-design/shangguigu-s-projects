package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {

	//定义一个存放任意数据类型的管道 3个数据
	//var allchan chan interface{}
	allchan := make(chan interface{}, 3)

	allchan <- 10
	allchan <- "tom jack"
	cat := Cat{"小花猫", 4}
	allchan <- cat

	//我们希望获得管道中的第三个元素，则将前两个推出
	<-allchan
	<-allchan

	newCat := <-allchan //从管道中取出的Cat是什么？

	fmt.Printf("newCat=%T,newCat=%v\n", newCat, newCat)
	//下面写法错误
	//fmt.Printf("newCat.Name=%v",newCat.Name)
	//使用类型断言
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v", a.Name)
}
