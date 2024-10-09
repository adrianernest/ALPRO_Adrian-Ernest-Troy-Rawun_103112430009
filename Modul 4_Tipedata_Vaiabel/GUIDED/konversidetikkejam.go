package main

import "fmt"


func main() {
	var jam, menit, detik, masukan int
	fmt.Scanln(&masukan)
	jam = masukan/3600
	menit = (masukan%3600)/60
	detik = masukan %60
	fmt.Println(jam, "jam", menit, "menit", detik, "detik")

}
