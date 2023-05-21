package main

import "fmt"

func main() {
	var first, second int
	result := false
	fmt.Scan(&first, &second)

	for i := second; i >= first; i-- {
		if i % 7 == 0 || i == 0 {
			fmt.Print(i)
			result = true
			break
		}
	}
	
	if !result {
		fmt.Print("NO")
	}
}
