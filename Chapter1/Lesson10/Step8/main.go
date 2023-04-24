package main

import "fmt"

func main () {
	var i int
	for fmt.Scan(&i); i <= 100; fmt.Scan(&i){
		if i < 10 {
			continue
		} else {
			fmt.Print(i)
		}
	}
}