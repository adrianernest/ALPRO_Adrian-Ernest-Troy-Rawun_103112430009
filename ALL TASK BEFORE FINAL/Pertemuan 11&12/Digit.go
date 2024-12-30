package main

import "fmt"

func main() {
	var X, sum int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scanln(&X)

	// Menampilkan digit dan menghitung jumlah digit
	for X > 0 {
		digit := X % 10      // Mendapatkan digit terakhir
		fmt.Print(digit, " ") // Menampilkan digit
		sum += digit         // Menambahkan digit ke jumlah
		X /= 10              // Menghapus digit terakhir
	}

	// Menampilkan jumlah total digit
	fmt.Println() // Untuk baris baru setelah digit
	fmt.Println("Total penjumlahan digit:", sum)
}
