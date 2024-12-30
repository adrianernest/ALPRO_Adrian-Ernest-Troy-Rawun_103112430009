package main

import "fmt"

func main() {
    var fisikSehat, perbekalanCukup, cuacaBagus bool

    fmt.Scan(&fisikSehat)

    fmt.Scan(&perbekalanCukup)

    fmt.Scan(&cuacaBagus)

    // Menentukan apakah mahasiswa akan berkemah
    if fisikSehat && perbekalanCukup && cuacaBagus {
        fmt.Println("berkemah")
    } else {
        fmt.Println("tidak berkemah")
    }
}
