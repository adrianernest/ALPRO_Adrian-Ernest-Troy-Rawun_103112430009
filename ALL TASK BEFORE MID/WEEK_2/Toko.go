package main

import "fmt"

func main() {
	var hargaBeli []int
	var hargaJual []float64

	fmt.Print("Masukkan tiga bilangan bulat positif: ")
	fmt.Scanln(&hargaBeli[0], &hargaBeli[1], &hargaBeli[2])

	for i := 0; i < len(hargaBeli); i++ {
		hargaJual[i] = float64(hargaBeli[i]) * 1.05
	}

	fmt.Println("Keluaran: ", hargaJual)
}
