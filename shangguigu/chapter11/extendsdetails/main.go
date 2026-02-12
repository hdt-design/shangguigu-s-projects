package main

import (
	"fmt"
)

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
	Name string
}

func (b *B) SayOk() {
	fmt.Println("B SayOk", b.Name)
}

func main() {

	// var b B
	// b.A.Name = "tom"
	// b.A.age = 19
	// b.A.SayOk() //A SayOk tom
	// b.A.hello() //A hello tom

	// //上面的写法可以简化

	// b.Name = "smith"
	// b.age = 20
	// b.SayOk() //B SayOk smith
	// b.hello() //A hello tom

	var b B
	b.Name = "jack"
	b.A.Name = "scott"
	b.age = 100
	b.SayOk()   //B SayOk jack
	b.A.SayOk() //A SayOk scott
	b.hello()   //A hello scott
}
