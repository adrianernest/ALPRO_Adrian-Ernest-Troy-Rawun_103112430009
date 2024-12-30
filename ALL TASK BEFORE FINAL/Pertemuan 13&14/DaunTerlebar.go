package main

import "fmt"

func main() {
    var n, max int
    fmt.Scan(&n) // Membaca jumlah daun

    for i := 0; i < n; i++ {
        var lebar int
        fmt.Scan(&lebar)
        if lebar > max {
            max = lebar // Memperbarui lebar maksimum
        }
    }

    fmt.Println(max) // Menampilkan lebar daun terbesar
}
