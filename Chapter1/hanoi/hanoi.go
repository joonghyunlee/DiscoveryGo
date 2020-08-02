package main

import "fmt"

func Hanoi(n int) {
    Helper(n, 1, 2, 3)
}

func Helper(n int, from int, to int, aux int) {
    if n == 1 {
        fmt.Println(from, "->", to)
        return
    }

    Helper(n - 1, from, aux, to)
    fmt.Println(from, "->", to)
    Helper(n - 1, aux, to, from)
}

func main() {
    Hanoi(3)
}
