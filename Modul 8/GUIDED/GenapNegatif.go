package main

import (
    "fmt"
)

func main() {
    var a int
    var b bool

    fmt.Print("Masukkan bilangan bulat: ")
    fmt.Scan(&a)

    b = a < 0 && a%2 == 0

    fmt.Println(b)
}
