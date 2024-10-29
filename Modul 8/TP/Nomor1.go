package main // Mendefinisikan paket utama untuk program

import "fmt" // Mengimpor paket fmt untuk input dan output

func main() {
	// Mendeklarasikan variabel untuk menyimpan nilai ujian
	var nilaiUjian int

	// Menampilkan pesan kepada pengguna untuk memasukkan nilai ujian
	fmt.Print("Masukkan nilai ujian: ")

	// Menerima input dari pengguna dan menyimpannya di variabel nilaiUjian
	fmt.Scan(&nilaiUjian)

	// Mengecek apakah nilai ujian lebih besar atau sama dengan 70
	if nilaiUjian >= 70 {
		// Jika kondisi benar, tampilkan "Lulus"
		fmt.Println("Lulus")
	} else {
		// Jika kondisi salah, tampilkan "Tidak Lulus"
		fmt.Println("Tidak Lulus")
	}
}
