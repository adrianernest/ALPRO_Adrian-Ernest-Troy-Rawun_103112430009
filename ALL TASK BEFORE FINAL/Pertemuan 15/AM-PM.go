package main

import "fmt"

func main() {
    var jam int
    fmt.Scan(&jam) // Membaca jam dalam format 24 jam

    if jam == 0 {
        fmt.Println("12 AM")
    } else if jam < 12 {
        fmt.Printf("%d AM\n", jam)
    } else if jam == 12 {
        fmt.Println("12 PM")
    } else {
        fmt.Printf("%d PM\n", jam-12) // Konversi ke format PM
    }
}
