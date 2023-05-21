package main

import "fmt"

func main() {
	var (
		number int
		fibCount = 2
		fibCurr = 2
		fibPrev = 1
	)

	fmt.Scan(&number)

	for fibCurr <= number {
		fibPrev, fibCurr = fibCurr, fibPrev + fibCurr
		fibCount++
	}

	if number == fibPrev {
		fmt.Print(fibCount)
	} else {
		fmt.Print(-1)
	}
}