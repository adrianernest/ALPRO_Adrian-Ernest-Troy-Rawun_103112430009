package main

import "fmt"

func main() {
    fmt.Println("Kalkulator Sederhana")
    fmt.Println("---------------------")

    var (
        num1 float64
        num2 float64
        operasi string
    )

    fmt.Print("Masukkan angka pertama: ")
    fmt.Scanln(&num1)

    fmt.Print("Masukkan operasi (+, -, *, /): ")
    fmt.Scanln(&operasi)

    fmt.Print("Masukkan angka kedua: ")
    fmt.Scanln(&num2)

    switch operasi {
    case "+":
        fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
    case "-":
        fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
    case "*":
        fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
    case "/":
        if num2 != 0 {
            fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
        } else {
            fmt.Println("Error: Pembagian dengan nol!")
        }
    default:
        fmt.Println("Error: Operasi tidak valid!")
    }
}