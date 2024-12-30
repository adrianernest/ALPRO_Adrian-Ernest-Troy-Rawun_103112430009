package main

import (
	"fmt"
)

func main() {
	var num, d, counter int

	// Input angka dari user
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&num)

	counter = 0

	// Loop untuk memproses setiap digit dalam bilangan
	for num > 0 {
		d = num % 10 // Mendapatkan digit terakhir
		if d%2 == 0 && d != 0 { // Cek apakah digit genap dan bukan nol
			counter++
		}
		num = num / 10 // Menghilangkan digit terakhir
	}

	// Output jumlah digit genap 
	fmt.Println("Jumlah digit genap :", counter)
}