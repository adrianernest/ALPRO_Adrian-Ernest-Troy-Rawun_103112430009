package main

import "fmt"

func main() {
    var x int
    fmt.Scan(&x) // Membaca bilangan

    for i := 0; i < x; i++ {
        for j := 1; j <= x; j++ {
            fmt.Print(j, " ")
        }
        fmt.Println() // Pindah ke baris berikutnya
    }
}
