package main

import (
	"fmt"
	"time"
)

// 向intChan中放入1-8000个数
func putNum(intChan chan int) {

	for i := 1; i <= 8000; i++ {
		intChan <- i
	}

	//关闭intChan
	close(intChan)
}

// 从intChan取出数据，并判断是否为素数，如果是，
// 就放入primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	//使用for循环
	//var num int
	var flag bool
	for {
		//time.Sleep(time.Millisecond*10)
		num, ok := <-intChan //intChan 取不到
		if !ok {
			break
		}
		flag = true //假定为素数
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //不是素数
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}

	fmt.Println("有一个primeNum协程因为取不到数据，退出")
	//这里还不能关闭primeChan
	//向exitChan写入true
	exitChan <- true
}

func main() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	//标识退出的管道
	exitChan := make(chan bool, 4) //4个

	start := time.Now().UnixMilli()

	//开启一个协程，向intChan中放入1-8000个数
	go putNum(intChan)
	//开启四个协程，从intChan取出数据，并判断是否为素数，如果是，
	//就放入primeChan

	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		end := time.Now().UnixMilli()
		fmt.Println("使用耗时", end-start)

		//取出四个结果就close
		close(primeChan)
	}()

	//遍历primeChan，结果输出
	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		//fmt.Println("素数=", res)
	}
	fmt.Println("main进程退出")
}
