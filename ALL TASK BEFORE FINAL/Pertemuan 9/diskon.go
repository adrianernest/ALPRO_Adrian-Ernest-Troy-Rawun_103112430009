package main

import "fmt"

func main() {
    var totalBelanja float64
    var sedangAsesmen bool
    fmt.Scan(&totalBelanja, &sedangAsesmen) // Membaca total belanja dan status asesmen
    if sedangAsesmen {
        totalBelanja *= 0.65 // Jika sedang asesmen, potong 35%
    }
    fmt.Println(totalBelanja) // Menampilkan total belanja setelah diskon (jika ada)
}
