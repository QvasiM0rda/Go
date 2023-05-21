package main

import "fmt"

func main() {
	var first, second, third int
	fmt.Scanf("%1d%1d%1d", &first, &second, &third)
	fmt.Printf("%d%d%d", third, second, first)
}