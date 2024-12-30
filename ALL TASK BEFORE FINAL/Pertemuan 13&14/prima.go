package main

import "fmt"

func main() {
    var x int
    fmt.Scan(&x) // Membaca bilangan input

    if x < 2 {
        fmt.Println(false) // Bilangan kurang dari 2 bukan prima
        return
    }

    isPrima := true
    for i := 2; i*i <= x; i++ {
        if x%i == 0 {
            isPrima = false // Jika habis dibagi selain 1 dan dirinya, bukan prima
            break
        }
    }

    fmt.Println(isPrima) // Menampilkan hasil apakah prima atau tidak
}
