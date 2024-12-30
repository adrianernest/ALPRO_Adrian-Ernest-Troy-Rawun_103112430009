package main

import "fmt"

func main() {
	var n, m, x, y int
	fmt.Print("Masukkan jumlah gula (n), kopi (m), gula per cangkir (x), dan kopi per cangkir (y): ")
	fmt.Scanln(&n, &m, &x, &y)

	var count int

	// Menghitung banyaknya cangkir kopi yang bisa dibuat
	for n >= x && m >= y {
		n -= x // Mengurangi gula sesuai dengan kebutuhan
		m -= y // Mengurangi kopi sesuai dengan kebutuhan
		count++ // Menambah jumlah cangkir
	}

	// Menampilkan jumlah cangkir kopi yang bisa dibuat
	fmt.Println("Banyaknya cangkir kopi yang bisa dibuat:", count)
}
