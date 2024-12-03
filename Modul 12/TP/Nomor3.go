package main

import "fmt"

func main() {
	var harga, totalBelanja int

	// Perulangan untuk meminta input harga barang satu per satu
	for {
		// Meminta input harga barang
		fmt.Print("Masukkan harga barang (ketik 0 untuk selesai): ")
		fmt.Scanln(&harga)

		// Jika harga yang dimasukkan adalah 0, keluar dari perulangan
		if harga == 0 {
			break
		}

		// Menambahkan harga barang ke total belanja
		totalBelanja += harga
	}

	// Menampilkan total belanja
	fmt.Printf("Total belanja Anda: %d\n", totalBelanja)
}
