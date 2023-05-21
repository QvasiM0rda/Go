package main

import "fmt"

func sumInt(arguments ...int) (int, int) {
	sum := 0
	for _, arg := range arguments {
		sum += arg
	}
	return len(arguments), sum
}

func main() {
	fmt.Print(sumInt(1, 0))
}