package main

import "fmt"

func main() {
	// Kamus
	var y, bilangan int
	var angka [9]int

	// Input nilai y
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	// Input sembilan bilangan
	fmt.Println("Masukkan 9 bilangan:")
	for i := 0; i < 9; i++ {
		fmt.Scan(&bilangan)
		angka[i] = bilangan
	}

	// Proses pengurutan menggunakan metode bubble sort
	for i := 0; i < 9; i++ {
		for j := 0; j < 8-i; j++ {
			if angka[j] > angka[j+1] {
				// Tukar posisi
				angka[j], angka[j+1] = angka[j+1], angka[j]
			}
		}
	}

	// Median adalah elemen ke-5 (indeks ke-4 setelah diurutkan)
	median := angka[4]

	// Output
	fmt.Printf("Median bernilai %d\n", median)
}
