package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)
	fmt.Printf("It is %d hours %d minutes.", k / 3600, k % 3600 / 60)
}