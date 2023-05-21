package main

import "fmt"

func main() {
	var num float64
	fmt.Scan(&num)
	switch {
	case num <= 0:
		fmt.Printf("число %2.2f не подходит", num)
	case num > 10000:
		fmt.Printf("%e", num)
	default:
		fmt.Printf("%.4f", num*num)
	}
}
