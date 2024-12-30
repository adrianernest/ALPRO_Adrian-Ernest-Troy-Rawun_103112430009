package main

import (
	"fmt"
)

func main() {
	// Input data
	var h1, m1, h2, m2 int

	fmt.Println("Masukkan jam dan menit mulai parkir (h1 m1):")
	fmt.Scan(&h1, &m1)
	fmt.Println("Masukkan jam dan menit selesai parkir (h2 m2):")
	fmt.Scan(&h2, &m2)

	// Konversi waktu ke menit
	startMinutes := h1*60 + m1
	endMinutes := h2*60 + m2

	// Validasi input
	if startMinutes < 7*60 || startMinutes > 17*60 || endMinutes < 7*60 || endMinutes > 17*60 || startMinutes > endMinutes {
		fmt.Println("Waktu parkir tidak valid. Periksa input Anda!")
		return
	}

	// Hitung durasi parkir
	duration := endMinutes - startMinutes
	hh := duration / 60
	mm := duration % 60

	// Format output
	fmt.Printf("Durasi parkir: %d jam %d menit\n", hh, mm)
	fmt.Printf("Parkir dari %02d:%02d sampai %02d:%02d\n", h1, m1, h2, m2)
}
