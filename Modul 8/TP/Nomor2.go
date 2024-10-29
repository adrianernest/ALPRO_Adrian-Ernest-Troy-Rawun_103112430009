package main  // Mendefinisikan paket utama untuk program

import "fmt"  // Mengimpor paket fmt untuk fungsi input dan output

func main() {
    // Mendeklarasikan variabel untuk menyimpan angka yang akan dimasukkan oleh pengguna
    var angka int
    
    // Menampilkan pesan kepada pengguna untuk memasukkan angka
    fmt.Print("Masukkan Angka: ")
    
    // Mengambil input dari pengguna dan menyimpannya di variabel angka
    fmt.Scan(&angka)

    // Memeriksa apakah angka tersebut habis dibagi 2
    if angka%2 == 0 {
        // Jika angka habis dibagi 2, tampilkan pesan "Angka adalah Genap"
        fmt.Println("Angka adalah Genap")
    } else {
        // Jika angka tidak habis dibagi 2, tampilkan pesan "Angka adalah Ganjil"
        fmt.Println("Angka adalah Ganjil")
    }
}
