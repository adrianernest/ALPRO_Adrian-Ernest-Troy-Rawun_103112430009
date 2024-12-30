package main

import "fmt"

func main (){
	var totalBelanjaAwal, totalBelanjaAkhir, diskon int

	fmt.Print("masukan total belanja awal:")
	fmt.Scanln(&totalBelanjaAwal)

	fmt.Print("masukan diskonnya:")
	fmt.Scanln(&diskon)

	totalBelanjaAkhir = totalBelanjaAwal-(totalBelanjaAwal*diskon/100)
	fmt.Print("total  belanja akhir: ", totalBelanjaAkhir)

}