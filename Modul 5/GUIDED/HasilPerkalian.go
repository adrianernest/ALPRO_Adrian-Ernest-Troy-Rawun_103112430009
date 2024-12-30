package main

import "fmt"

func main() {
    var bil1, bil2, hasil int

    // Meminta input dua bilangan
    fmt.Print("Masukkan bilangan pertama: ")
    fmt.Scan(&bil1)
    fmt.Print("Masukkan bilangan kedua: ")
    fmt.Scan(&bil2)

    // Melakukan perkalian tanpa operator *
    for i := 0; i < bil2; i++ {
        hasil += bil1 // Menambahkan bil1 ke hasil sebanyak bil2 kali
    }

    // Menampilkan hasil perkalian
    fmt.Printf("Hasil perkalian %d dan %d adalah: %d\n", bil1, bil2, hasil)
}
