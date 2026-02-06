package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

type Rect struct {
	leftup, rightdown Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	//r1有四个int，在内存中连续分布
	//打印地址
	fmt.Printf("r1.leftup.X 地址=%p\n", &r1.leftup.X)
	fmt.Printf("r1.leftup.Y 地址=%p\n", &r1.leftup.Y)
	fmt.Printf("r1.rightdown.X 地址=%p\n", &r1.rightdown.X)
	fmt.Printf("r1.rightdown.Y 地址=%p\n", &r1.rightdown.Y)

	//output
	// 	r1.leftup.X 地址=0x1400001c060
	// r1.leftup.Y 地址=0x1400001c068
	// r1.rightdown.X 地址=0x1400001c070
	// r1.rightdown.Y 地址=0x1400001c078

	
}
