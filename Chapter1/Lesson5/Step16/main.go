package main

import "fmt"

func main() {
	var d int
	fmt.Scan(&d)
	fmt.Println("It's", d/30, "hours", 2*(d%30), "minutes.")
}
