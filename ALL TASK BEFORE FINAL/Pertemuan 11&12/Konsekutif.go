package main

import "fmt"

func main() {
	var X int
	fmt.Print("Masukkan bilangan bulat positif X: ")
	fmt.Scanln(&X)

	// Ubah X menjadi string untuk memudahkan pengambilan digit
	strX := fmt.Sprint(X)

	// Periksa setiap pasangan digit berturut-turut
	for i := 0; i < len(strX)-1; i++ {
		// Menghitung selisih antara dua digit berturut-turut
		if abs(int(strX[i])-int(strX[i+1])) != 1 {
			fmt.Println("Bilangan bukan konsekutif")
			return
		}
	}

	// Jika semua digit berurutan selisihnya 1
	fmt.Println("Bilangan adalah konsekutif")
}

// Fungsi untuk menghitung nilai absolut
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
