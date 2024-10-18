package main

import "fmt"

func main() {
    var a, b int
    fmt.Print("Masukan Nilai  A: ")

    fmt.Scan(&a)
    fmt.Print("Masukan Nilai   B: ")

    fmt.Scan(&b)

    if a <= b {
        for i := a; i <= b; i++ {
            fmt.Print(i, " ")
        }
    } else {
        fmt.Println("Input tidak  valid")

    }
}