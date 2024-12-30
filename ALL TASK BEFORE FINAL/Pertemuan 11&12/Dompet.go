package main

import "fmt"

func main() {
	var saldo, transaksi int

	// Loop untuk menerima transaksi
	for {
		fmt.Print("Masukkan transaksi (0 untuk selesai): ")
		fmt.Scanln(&transaksi)

		if transaksi == 0 {
			break // Akhiri loop jika transaksi adalah 0
		}

		saldo += transaksi // Tambahkan transaksi ke saldo
	}

	// Tampilkan saldo akhir
	fmt.Printf("Saldo akhir dalam dompet: %d\n", saldo)
}
