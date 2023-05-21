package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5, 5)
	a[1] = 1
	fmt.Print(a)
}