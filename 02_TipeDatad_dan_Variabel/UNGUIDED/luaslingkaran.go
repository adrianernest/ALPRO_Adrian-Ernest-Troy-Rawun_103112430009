package main

import "fmt"

func hitungLuasLingkaran(r float64) float64 {
  const phi = 3.14
  return phi * r * r
}

func main() {
  var r float64

  fmt.Print("Masukkan jari-jari lingkaran: ")
  fmt.Scanln(&r)

  if r < 0 {
    fmt.Println("Jari-jari tidak boleh negatif")
  } else {
    luas := hitungLuasLingkaran(r)
    fmt.Printf("Luas lingkaran dengan jari-jari %.2f adalah %.2f\n", r, luas)
  }
}