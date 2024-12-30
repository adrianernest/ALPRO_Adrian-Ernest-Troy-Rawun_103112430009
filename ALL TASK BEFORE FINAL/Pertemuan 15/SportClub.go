package main

import "fmt"

func main() {
    var s int
    fmt.Scan(&s) // Membaca jumlah jatah renang

    minggu := s / 7 // Hitung minggu penuh
    if s%7 != 0 {
        minggu++ // Tambahkan minggu jika ada sisa hari
    }

    fmt.Printf("Minggu ke-%d\n", minggu)
}
