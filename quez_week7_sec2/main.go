package main

import "fmt"

func printNumbers(n int) {
    if n < 1 {
        return
    }
    fmt.Println(n)
    printNumbers(n - 1)
}

func main() {
    printNumbers(10)
}