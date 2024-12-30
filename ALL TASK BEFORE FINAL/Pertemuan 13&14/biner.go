package main

import "fmt"

func main() {
    var x int
    fmt.Scan(&x) // Membaca bilangan desimal

    var biner string
    for x > 0 {
        biner = fmt.Sprintf("%d", x%2) + biner // Menambahkan sisa pembagian ke hasil
        x /= 2
    }

    fmt.Println(biner) // Menampilkan hasil konversi ke biner
}
