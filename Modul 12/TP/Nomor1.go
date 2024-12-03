package main

import "fmt"

func main() {
	// Menentukan angka rahasia
	var angkaRahasia = 7 // Anda bisa mengubah angka ini sesuai keinginan

	var tebakan int

	// Simulasi repeat...until
	for {
		// Meminta input dari pengguna
		fmt.Print("Tebak angka (1-10): ")
		_, err := fmt.Scan(&tebakan)
		if err != nil {
			fmt.Println("Input tidak valid, coba lagi.")
			// Jika input tidak valid, program akan mengulang
			continue
		}

		// Memeriksa apakah tebakan sesuai dengan angka rahasia
		if tebakan == angkaRahasia {
			fmt.Println("Selamat, tebakan Anda benar!")
			break // Keluar dari perulangan setelah tebakan benar
		} else {
			fmt.Println("Tebakan Anda salah, coba lagi.")
		}
	}
}
