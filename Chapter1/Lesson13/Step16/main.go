package main

import "fmt"

func main() {
	var (
		number int
		digit int
		slice []int
	)
	fmt.Scan(&number, &digit)

	for number > 0 {
		if number % 10 != digit {
			slice = append(slice, number % 10)
		}
		number /= 10
	}

	for i := len(slice) - 1; i >= 0; i-- {
		fmt.Print(slice[i])
	}
}