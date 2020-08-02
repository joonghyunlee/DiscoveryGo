package main

import "fmt"

func fibonacci(n int) int {
    p, q := 0, 1

    for n - 1 > 0 {
        p, q = q, p + q
        n--
    }

    return p
}

func main() {
    var n = fibonacci(7)
    fmt.Println(n)
}
