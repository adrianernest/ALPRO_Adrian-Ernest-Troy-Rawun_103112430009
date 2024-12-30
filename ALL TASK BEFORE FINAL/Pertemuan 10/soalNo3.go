package main

import (
	"fmt"
)

func main() {
	// Input keuntungan bulan ini dan bulan sebelumnya
	var bulanSebelumnya, bulanIni float64

	fmt.Println("Masukkan keuntungan bulan sebelumnya:")
	fmt.Scan(&bulanSebelumnya)
	fmt.Println("Masukkan keuntungan bulan ini:")
	fmt.Scan(&bulanIni)

	// Hitung selisih keuntungan
	selisih := bulanIni - bulanSebelumnya

	// Tentukan output berdasarkan selisih
	if selisih > 0 {
		fmt.Printf("Peningkatan sebesar %.2f\n", selisih)
	} else if selisih < 0 {
		fmt.Printf("Penurunan sebesar %.2f\n", -selisih)
	} else {
		fmt.Println("Tetap")
	}
}
