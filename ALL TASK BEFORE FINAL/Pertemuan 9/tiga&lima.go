package main

import "fmt"

func main() {
    var num int
    fmt.Scan(&num) // Membaca bilangan
    if num%3 == 0 {
        // Jika kelipatan 3
        fmt.Println("Kelipatan 3")
    }
    if num%5 == 0 {
        // Jika kelipatan 5
        fmt.Println("Kelipatan 5")
    }
    // Tidak menampilkan apa pun jika bukan kelipatan 3 atau 5
}
