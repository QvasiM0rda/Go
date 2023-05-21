package main

import "fmt"

func main() {
	var n, x int
	count := 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x == 0 {
			count++
		}
	}
	fmt.Print(count)
}