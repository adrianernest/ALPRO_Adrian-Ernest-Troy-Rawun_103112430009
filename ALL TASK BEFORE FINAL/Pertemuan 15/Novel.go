package main

import "fmt"

func main() {
    var p, r int
    fmt.Scan(&p, &r) // Membaca jumlah halaman dan jumlah maksimal per hari

    hari := p / r // Hitung hari penuh
    if p%r != 0 {
        hari++ // Tambahkan hari jika ada sisa halaman
    }

    fmt.Printf("%d hari\n", hari)
}
