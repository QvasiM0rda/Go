package main

import "fmt"

func vote(x, y, z int) int {
	if x + y + z > 1 {
		return 1
	}
	return 0
}

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Print(vote(x, y, z))
}