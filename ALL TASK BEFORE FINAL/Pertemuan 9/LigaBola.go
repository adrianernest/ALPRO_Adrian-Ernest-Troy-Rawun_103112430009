package main

import "fmt"

func main() {
    var a, b, c, d int
    fmt.Scan(&a, &b, &c, &d) // Membaca jumlah gol dari empat tim

    // Menentukan jumlah gol terbesar
    max := a
    if b > max {
        max = b
    }
    if c > max {
        max = c
    }
    if d > max {
        max = d
    }

    // Menentukan jumlah gol terkecil
    min := a
    if b < min {
        min = b
    }
    if c < min {
        min = c
    }
    if d < min {
        min = d
    }

    // Menampilkan jumlah gol terbesar dan terkecil
    fmt.Println(max, min)
}
