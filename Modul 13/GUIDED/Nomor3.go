package main

import "fmt"

func main() {
    var n int

    fmt.Print("Masukkan bilangan bulat positif: ")
    fmt.Scan(&n)

    fmt.Printf("Faktor-faktor dari %d adalah: ", n)
    
    for i := 1; i <= n; i++ {
        if n % i == 0 {
            fmt.Print(i, " ")
        }
    }
    fmt.Println()
}
