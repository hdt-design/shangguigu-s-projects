package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func text01(p Person) {
	fmt.Println(p.Name)
}

func text02(p *Person) {
	fmt.Println(p.Name)
}

func (p Person) text03() {
	p.Name = "jack"
	fmt.Println("text03()=", p.Name)
}

func (p *Person) text04() {
	p.Name = "tom"
	fmt.Println("text04()=", p.Name)
}

func main() {
	p := Person{"Alice"}
	text01(p)
	text02(&p)
	p.text03()                            //jack
	fmt.Println("main() p.Name=", p.Name) //Alice
	(&p).text03()                         //从形式上看传入地址，本质仍是值拷贝
	fmt.Println("main() p.Name=", p.Name) //Alice
	p.text04()                            //等价于（*p）text04()
	//tom
	fmt.Println("main() p.Name=", p.Name) //tom
}

// 1. 新增text04()方法，修改text03()方法，使其修改Person的Name字段。
// 2. 调用text03()方法时，传入值拷贝，因此不会影响到原对象。
// 3. 调用(&p).text03()方法时，表面传入地址，底层仍是值拷贝，因此会不影响到原对象。
// 4. 调用p.text04()方法时，表面值拷贝，本质上地址传入，因此会影响到原对象。
// 5. 调用(&p).text04()方法时，本质上地址传入，因此会影响到原对象。
