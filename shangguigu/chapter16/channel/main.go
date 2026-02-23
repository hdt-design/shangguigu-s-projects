package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age  int
}

func main() {

	//使用channel存放int
	var intChan chan int
	intChan = make(chan int, 3)
	intChan <- 10
	intChan <- 20
	intChan <- 30
	//继续存放会报错
	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan
	fmt.Printf("num1=%d,num2=%d,num3=%d\n", num1, num2, num3)

	/*-------------------------------------------------------------------------*/

	//使用channel存放map类型
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	m1 := make(map[string]string, 20)
	m1["city1"] = "北京"
	m1["city2"] = "天津"

	m2 := make(map[string]string, 20)
	m2["hero1"] = "宋江"
	m2["hero2"] = "武松"
	//...
	mapChan <- m1
	mapChan <- m2
	fmt.Println("mapChan=", mapChan)

	/*-------------------------------------------------------------------------*/

	//使用channel存放结构体类型
	var catChan chan Cat
	catChan = make(chan Cat, 10)

	cat1 := Cat{Name: "tom", Age: 18}
	cat2 := Cat{Name: "tom~", Age: 180}
	catChan <- cat1
	catChan <- cat2

	fmt.Println("catChan=", catChan)

	//取出
	cat11 := <-catChan
	cat22 := <-catChan
	fmt.Println("cat11=", cat11, "cat22=", cat22)

	/*-----------------------------------------------------------------------*/

	//存放结构体指针
	var catChan2 chan *Cat
	catChan2 = make(chan *Cat, 10)

	cat1 = Cat{Name: "tom", Age: 18}
	cat2 = Cat{Name: "tom~", Age: 180}
	catChan2 <- &cat1
	catChan2 <- &cat2

	cat3 := <-catChan2
	cat4 := <-catChan2
	fmt.Println("catChan2=", catChan2)
	fmt.Println(cat3, cat4)
}
