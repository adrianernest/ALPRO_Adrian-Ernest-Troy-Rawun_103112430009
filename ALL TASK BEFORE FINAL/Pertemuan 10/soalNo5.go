package main

import (
	"fmt"
)

func main() {
	// Input data
	var totalBelanja int
	var bersediaDibuatKartu bool

	fmt.Println("Masukkan total belanja:")
	fmt.Scan(&totalBelanja)
	fmt.Println("Apakah bersedia dibuatkan kartu? (true/false):")
	fmt.Scan(&bersediaDibuatKartu)

	// Logika promo
	mendapatKartu := bersediaDibuatKartu && totalBelanja >= 100000
	mendapatDiskon := totalBelanja >= 100000
	mendapatCashback := totalBelanja >= 200000

	// Perhitungan total setelah promo
	totalSetelahDiskon := totalBelanja
	if mendapatDiskon {
		totalSetelahDiskon = totalBelanja - (totalBelanja * 10 / 100)
	}

	totalSetelahCashback := totalSetelahDiskon
	if mendapatCashback {
		totalSetelahCashback -= 75000
	}

	// Output
	fmt.Println("Kartu?", mendapatKartu)
	fmt.Println("Diskon?", mendapatDiskon)
	fmt.Println("Cashback?", mendapatCashback)
	fmt.Printf("Rp. %d\n", totalSetelahCashback)
}
