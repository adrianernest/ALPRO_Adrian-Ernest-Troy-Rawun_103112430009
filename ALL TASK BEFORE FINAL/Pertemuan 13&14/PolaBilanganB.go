package main

import "fmt"

func main() {
    var x int
    fmt.Scan(&x) // Membaca bilangan

    for i := 1; i <= x; i++ {
        for j := 1; j <= x; j++ {
            if i == 1 || i == x || j == 1 || j == x || i == j || j == x-i+1 {
                fmt.Print(i, " ") // Menampilkan angka pada posisi tertentu
            } else {
                fmt.Print("  ") // Spasi kosong
            }
        }
        fmt.Println() // Pindah ke baris berikutnya
    }
}
