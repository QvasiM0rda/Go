package main

import "fmt"

func main() {
	var n, max, count int
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		switch {
		case max < n:
			max = n
			count = 1
		case max == n:
			count += 1
		}
	}
	fmt.Print(count)
}