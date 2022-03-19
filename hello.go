package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

// 型省略
// func add(x , y int) int {
// 	return x + y
// }

func swap(x, y string) (string, string) {
	return x, y
}

func main() {
	// fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
