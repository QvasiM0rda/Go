package main

import "fmt"

func main() {
	var x, p, y int
	fmt.Scan(&x, &p, &y)
	years := 1

	for {
		x += x * p / 100
		if x > y {
			break
		}
		years++
	}
	fmt.Print(years)
}