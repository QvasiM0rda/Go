package main

import "fmt"

func main() {
	workArray := [10]uint8{}
	for i := range workArray {
		fmt.Scan(&workArray[i])
	}

	var first, second uint8
	for i := 0; i < 3; i++ {
		fmt.Scan(&first, &second)
		workArray[first], workArray[second] = workArray[second], workArray[first]
	}

	for _, elem := range workArray {
		fmt.Print(elem, " ")
	}
}