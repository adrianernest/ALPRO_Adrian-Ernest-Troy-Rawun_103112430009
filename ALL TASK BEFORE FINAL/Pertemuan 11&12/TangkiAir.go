package main

import "fmt"

func main() {
	var T, V, totalAir int
	fmt.Print("Masukkan kapasitas tangki T: ")
	fmt.Scanln(&T)

	// Mengisi tangki hingga penuh
	for {
		fmt.Print("Masukkan volume air dari ember (0 untuk berhenti): ")
		fmt.Scanln(&V)
		if V == 0 {
			break // Jika volume ember 0, berhenti mengisi
		}

		totalAir += V // Menambahkan volume air ke tangki

		// Cek apakah tangki sudah penuh
		if totalAir >= T {
			fmt.Println("true") // Tangki penuh
		} else {
			fmt.Println("false") // Tangki belum penuh
		}
	}
}
