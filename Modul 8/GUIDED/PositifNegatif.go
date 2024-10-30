package main

import "fmt"

func main() {
    var bilangan int

    fmt.Print("Masukkan bilangan bulat: ")
    fmt.Scan(&bilangan)

    if bilangan <= 0 {
        fmt.Print("bukan positif")
    } else {
		fmt.Print("Positif")
	}
}