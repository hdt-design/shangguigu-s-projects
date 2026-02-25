package main

import (
	"fmt"
	"reflect"
)

// 专门演示反射
func reflectTest01(b interface{}) {

	//通过反射获取传入的变量type,kind，值
	//1.先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)

	n2 := 2 + rVal.Int()
	//n3 := rVal.Float()
	fmt.Println("n2=", n2)
	//fmt.Println("n3=", n3)

	fmt.Printf("rVal=%v,rVal type=%T\n", rVal, rVal)

	//下面我们将rVal转成interface{}
	iv := rVal.Interface()
	//将interface{}通过断言转成需要的类型
	num2 := iv.(int)
	fmt.Println("num2=", num2)
}

// 专门演示反射【对结构体的反射】
func reflectTest02(b interface{}) {

	//通过反射获取传入的变量的type,kind,值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rtyp=", rTyp)

	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)

	//3.获取变量对应的kind
	//（1）rVal.Kind()==>
	kind1 := rVal.Kind()
	//(2)rTyp.Kind()==>
	kind2 := rTyp.Kind()
	fmt.Printf("kind1=%v, kind=%v\n", kind1, kind2)

	//下面将rVal转化成interface{}
	iv := rVal.Interface()
	fmt.Printf("iv=%d iv type=%T", iv, iv)
	//将 interface{} 通过断言转成需要的类型
	//这里，我们就简单使用了一带检测的类型断言.
	//同学们可以使用 switch 的断言形式来做的更加的灵活
	stu, ok := iv.(Student)
	if ok {
		fmt.Println("stu=", stu)
	}
}

type Student struct {
	Name string
	Age  int
}

func main() {

	//请编写一个案例
	//演示对（基本数据类型，interface{},reflect.value）进行反射的基本操作

	//1.先定义一个int
	var num = 100
	reflectTest01(num)

	//2.定义一个Student实例
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)
}
