package main

import (
	"fmt"
)

func main() {
	var usia int

	fmt.Print("Masukkan usia Anda: ")
	fmt.Scan(&usia)

	if usia >= 0 && usia <= 12 {
		fmt.Println("Kategori: Anak-anak")
	} else if usia >= 13 && usia <= 17 {
		fmt.Println("Kategori: Remaja")
	} else if usia >= 18 && usia <= 64 {
		fmt.Println("Kategori: Dewasa")
	} else if usia >= 65 {
		fmt.Println("Kategori: Lansia")
	} else {
		fmt.Println("Usia tidak valid")
	}
}
