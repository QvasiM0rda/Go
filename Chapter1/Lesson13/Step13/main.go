package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	powers_of_two := 1
	for powers_of_two <= n {
		fmt.Print(powers_of_two, " ")
		powers_of_two *= 2
	}
}