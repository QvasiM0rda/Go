package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	left := (num / 100000) + (num % 100000 / 10000) + (num % 10000 / 1000)
	right := (num % 1000 / 100) + (num % 100 / 10) + num % 10
	if left == right {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}