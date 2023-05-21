package main

import "fmt"
import "time"

func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonacci(n - 1) + fibonacci(n - 2)
}

func main() {
	var n int
	fmt.Scan(&n)
	start := time.Now()
	fmt.Println(fibonacci(n))
	duration := time.Since(start)
	fmt.Println(duration)
}