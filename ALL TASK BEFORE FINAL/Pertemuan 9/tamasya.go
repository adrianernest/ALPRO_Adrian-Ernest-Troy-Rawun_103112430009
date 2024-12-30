package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n) // Membaca jumlah mahasiswa
    mobil := (n + 6) / 7 // Membulatkan ke atas, karena satu mobil muat 7 orang
    fmt.Println(mobil) // Menampilkan jumlah mobil yang dibutuhkan
}
