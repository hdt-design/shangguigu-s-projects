package main

import (
	"fmt"
)

func main() {

	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan) //close
	//close后不能再写入数到channel
	//intChan<-300
	fmt.Println("ok ok~")
	//管道关闭后，读取数据是可以的
	n1 := <-intChan
	fmt.Println("n1=", n1)

	//遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2 //放入100个数据到管道
	}

	//遍历管道不能用普通的for循环
	//for i:=0;i<len(intChan2);i++{

	//}
	//在遍历时，如果channel没有关闭，则会deadlock
	//如果channel已经关闭，则正常遍历，遍历结束退出遍历
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}
