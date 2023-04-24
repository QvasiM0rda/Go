package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	hundreds := num / 100
	tens := num % 100 / 10
	ones := num / 10
	if hundreds != tens && tens != ones && hundreds != ones {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}