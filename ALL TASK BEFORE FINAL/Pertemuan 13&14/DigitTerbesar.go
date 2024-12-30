package main

import "fmt"

func main() {
    var x int
    fmt.Scan(&x) // Membaca bilangan

    maxDigit := 0
    for x > 0 {
        digit := x % 10       // Mengambil digit terakhir
        if digit > maxDigit { // Memperbarui digit terbesar
            maxDigit = digit
        }
        x /= 10 // Menghapus digit terakhir
    }

    fmt.Println(maxDigit) // Menampilkan digit terbesar
}
