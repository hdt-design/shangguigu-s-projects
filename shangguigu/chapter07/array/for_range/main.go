package main

import (
	"fmt"
)

func main() {
	var heroes = [3]string{"宋江", "卢俊义", "吴用"}

	for i, v := range heroes {
		fmt.Printf("i=%v v=%v\n", i, v)
		fmt.Printf("heroes[%d]=%v\n", i, heroes[i])
	}
}
