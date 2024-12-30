package main

import (
	"fmt"
)

func main() {
	// Input data
	var t1, t2, t3, t4, t5 float64

	fmt.Println("Masukkan 5 catatan temperatur (pisahkan dengan spasi):")
	fmt.Scan(&t1, &t2, &t3, &t4, &t5)

	// Periksa pola perubahan temperatur
	if (t2 > t1 && t3 > t2 && t4 > t3 && t5 > t4) || (t2 < t1 && t3 < t2 && t4 < t3 && t5 < t4) {
		fmt.Println("Stabil naik/turun")
	} else {
		fmt.Println("Tidak stabil")
	}
}
