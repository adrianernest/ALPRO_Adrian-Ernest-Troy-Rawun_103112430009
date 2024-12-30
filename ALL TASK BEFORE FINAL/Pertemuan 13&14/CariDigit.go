package main

import (
    "fmt"
    "strings"
)

func main() {
    var x, n int
    fmt.Scan(&x, &n) // Membaca x dan n

    strX := fmt.Sprintf("%d", x) // Konversi ke string
    strN := fmt.Sprintf("%d", n) // Konversi ke string

    fmt.Println(strings.Contains(strN, strX)) // Mengecek apakah x ada di n
}
