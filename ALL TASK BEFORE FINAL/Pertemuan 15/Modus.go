package main

import "fmt"

func main() {
    var x, countX, count0 int
    fmt.Scan(&x) // Membaca nilai x

    for i := 0; i < 9; i++ {
        var bil int
        fmt.Scan(&bil) // Membaca 9 bilangan
        if bil == x {
            countX++
        } else {
            count0++
        }
    }

    if countX > count0 {
        fmt.Printf("Modus = %d\n", x)
    } else {
        fmt.Println("Modus = 0")
    }
}
