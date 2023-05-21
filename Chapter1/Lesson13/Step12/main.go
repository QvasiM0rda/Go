package main

import "fmt"

func main() {
    var n int
	fmt.Scan(&n)
	cow := "korov"
	ending := ""

	switch {
		case (n % 10 >= 5 && n % 10 <= 9) || (n >= 5 && n <= 20): ending = ""
		case n % 10 == 1: ending = "a"
		case n % 10 >= 2 && n % 10 <= 4: ending = "y"
	}
	fmt.Println(n, cow + ending)
}