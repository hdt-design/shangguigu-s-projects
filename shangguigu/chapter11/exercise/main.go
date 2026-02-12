package main

import "fmt"

type Usb interface {
	Say()
}
type Stu struct {
}

func (stu *Stu) Say() {
	fmt.Println("Say()")
}
func main() {
	var stu = Stu{}
	//var u Usb = stu
	// 错误！ 会报 Stu类型没有实现Usb接口 ,
	// 如果希望通过编译,  var u Usb = &stu
	var u Usb = &stu
	u.Say()
	fmt.Println("here", u)
}
