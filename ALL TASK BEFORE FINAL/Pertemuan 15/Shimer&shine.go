package main

import "fmt"

func main() {
    var tk1, tk2 int
    fmt.Scan(&tk1, &tk2) // Membaca intensitas dua tongkat

    // Memeriksa apakah satu genap dan satu ganjil
    if (tk1%2 == 0 && tk2%2 != 0) || (tk1%2 != 0 && tk2%2 == 0) {
        fmt.Println("berhasil") // Jika iya, bola salju bisa dibuat
    }
}
