package main

import "fmt"

func main() {
	var A, B, sum int
	fmt.Scan(&A, &B)

	for i := A; i <= B; i++ {
		sum += i
	}
	fmt.Print(sum)
}