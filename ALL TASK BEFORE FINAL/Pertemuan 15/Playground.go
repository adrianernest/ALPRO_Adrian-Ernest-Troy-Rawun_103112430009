package main

import "fmt"

func main() {
    var member string
    var durasi int
    fmt.Scan(&member, &durasi) // Membaca jenis membership dan durasi bermain

    tarif := 65000 * 2 // Tarif 2 jam pertama
    if durasi > 2 {
        tarif += 20000 * (durasi - 2) // Tambahkan tarif kelebihan jam
    }

    if member == "Gold" {
        tarif /= 2 // Diskon 50% untuk Gold
    } else if member == "Silver" {
        tarif = tarif * 75 / 100 // Diskon 25% untuk Silver
    }

    fmt.Printf("IDR %d\n", tarif)
}
