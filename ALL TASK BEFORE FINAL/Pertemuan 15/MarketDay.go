package main

import "fmt"

func main() {
    var x, y, z int
    fmt.Scan(&x, &y, &z) // Membaca diskon (x), maksimum (y), dan total belanja (z)

    diskon := z * x / 100 // Hitung diskon
    if diskon > y {
        diskon = y // Jika diskon lebih besar dari maksimum, gunakan maksimum
    }

    fmt.Printf("%d rupiah\n", z-diskon) // Hitung total belanja akhir
}
