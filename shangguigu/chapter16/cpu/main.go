package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)

	//可以自己设置cpu
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
