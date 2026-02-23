package main

import (
	"fmt"
	"sync"
	"time"
)

// 需求：现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到map中。
// 最后显示出来。要求使用goroutine完成

// 思路
// 1. 编写一个函数，来计算各个数的阶乘，并放入到 map中.
// 2. 我们启动的协程多个，统计的将结果放入到 map中
// 3. map 应该做出一个全局的.

var (
	mymap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	//lock是一个全局互斥锁
	//sync是包：synchronized同步
	//mutex:是互斥
	lock sync.Mutex
)

// test函数计算n!，将这个结果昂如myMap
func test(n int) {

	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	//将res放入myMap
	//加锁
	lock.Lock()
	mymap[n] = res
	lock.Unlock()
}

func main() {

	//开启多个协程
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	//休眠10秒
	time.Sleep(10 * time.Second)

	//输出结果
	lock.Lock()
	for i, v := range mymap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
