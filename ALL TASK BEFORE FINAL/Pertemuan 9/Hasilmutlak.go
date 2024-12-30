package main 

import "fmt"

func main () {
    var bilangan int  
    fmt.Print("Masukkan bilangan: ")
    fmt.Scanln(&bilangan)

    if bilangan < 0 {
        bilangan = -1 * bilangan
    }
    fmt.Println("hasil mutlaknya adalah:", bilangan)
}