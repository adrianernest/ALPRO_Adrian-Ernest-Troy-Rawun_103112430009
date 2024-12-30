package main

import "fmt"

func main() {
    var kamar, bermain, kebun int
    fmt.Scan(&kamar, &bermain, &kebun) // Membaca persentase kebersihan

    if (kamar >= 60 && bermain >= 75 && kebun >= 60) || (kamar >= 80 && kebun >= 80) || (kamar == 100) {
        fmt.Println("Ice Cream")
    } else {
        fmt.Println("Tidak")
    }
}
