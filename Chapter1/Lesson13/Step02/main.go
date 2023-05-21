package main

import "fmt"

func main() {
	var (
		ones, tens, hundreds int
	)
	fmt.Scanf("%1d%1d%1d", &ones, &tens, &hundreds)
	fmt.Print(ones + tens + hundreds)
}