package main

import "fmt"

func minimumFromFour() int {
    var min, next int
    fmt.Scan(&min)
    for i := 0; i < 3; i++ {
        fmt.Scan(&next)
        if next < min {
            min = next
        }
    }
    return min
}

func main() {
	fmt.Print(minimumFromFour())
}