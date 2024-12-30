package main

import "fmt"

func main() {
    var tipe string
    countA, countB, countC := 0, 0, 0

    for {
        fmt.Scan(&tipe)
        if tipe != "A" && tipe != "B" && tipe != "C" {
            break // Berhenti jika input bukan A, B, atau C
        }

        switch tipe {
        case "A":
            countA++
        case "B":
            countB++
        case "C":
            countC++
        }
    }

    fmt.Println("Tipe A =", countA)
    fmt.Println("Tipe B =", countB)
    fmt.Println("Tipe C =", countC)
}
