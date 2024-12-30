package main

import (
	"fmt"
)

func main() {
	// Input sisi-sisi segitiga
	var a, b, c int
	fmt.Println("Masukkan tiga sisi segitiga (pisahkan dengan spasi):")
	fmt.Scan(&a, &b, &c)

	// Logika untuk menentukan jenis segitiga
	if a == b && b == c {
		fmt.Println("Segitiga Sama Sisi")
	} else if a == b || b == c || a == c {
		fmt.Println("Segitiga Sama Kaki")
	} else {
		fmt.Println("Segitiga Sembarang")
	}
}
