package main

import "fmt"

func main() {
    var jam, menit, durasi int
    fmt.Scan(&jam, &menit, &durasi) // Membaca jam, menit, dan durasi

    selesai := jam*60 + menit + durasi // Konversi ke menit dan tambahkan durasi
    jamSelesai := selesai / 60
    menitSelesai := selesai % 60

    if jamSelesai >= 8 && jamSelesai < 20 { // Cek apakah selesai sebelum 20:00
        fmt.Printf("%02d %02d\n", jamSelesai, menitSelesai)
    } else {
        fmt.Println("Silakan datang kembali besok")
    }
}
