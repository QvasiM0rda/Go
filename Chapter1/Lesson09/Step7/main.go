package main

import "fmt"

func main() {
    var num int
    fmt.Scan(&num)
    switch {
        case num < 10: fmt.Print(num)
        case num < 100: fmt.Print(num / 10)
        case num < 1000: fmt.Print(num / 100)
        case num < 10000: fmt.Print(num / 1000)
        default: fmt.Print(num / 10000)
    }
}