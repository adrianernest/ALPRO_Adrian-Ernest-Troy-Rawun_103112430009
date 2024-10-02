package main

import "fmt"

func main() {
    // Panjang sisi alun alun purwokerto
    sisi := 27.0

    // Menghitung keliling persegi
    keliling := 4 * sisi

    // Menghitung luas persegi
    luas := sisi * sisi

    // Menampilkan hasil
    fmt.Printf("Panjang sisi alun alun purwokerto: %.2f meter\n", sisi)
    fmt.Printf("Keliling alun alun purwokerto: %.2f meter\n", keliling)
    fmt.Printf("Luas alun alun purwokerto: %.2f meter persegi\n", luas)
}