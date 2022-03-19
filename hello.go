package main

import (
	"fmt"
)

func split(some int) (x, y int) {
	x = some * 4 / 9
	y = some - x
	return
}

func main() {
	fmt.Println(split(17))
}
