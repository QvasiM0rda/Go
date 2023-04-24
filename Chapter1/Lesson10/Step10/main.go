package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b)
	d = 1

	for ; a > 0; a /= 10{
		second := b
		first := a % 10
		for ; second > 0; second /= 10 {
			if first == second % 10 {
				c += first * d
				d *= 10
				continue
			}
		}
	}
	fmt.Print(c)
}