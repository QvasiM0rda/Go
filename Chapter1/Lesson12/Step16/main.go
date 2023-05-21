package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	array := make([]int, n, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}
	positive := 0

	for i := 0; i < n; i++ {
		if array[i] > 0 {
			positive += 1
		}
	}

	fmt.Print(positive)
}
