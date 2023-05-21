package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	array := make([]int, n, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	for i := 0; i < n; i += 2 {
		fmt.Print(array[i], " ")
	}
}
