package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if (a * a + b * b == c * c) || (a * a == b * b + c * c) || (a * a + c * c == b * b) {
		fmt.Print("Прямоугольный")
	} else {
		fmt.Print("Непрямоугольный")
	}
}