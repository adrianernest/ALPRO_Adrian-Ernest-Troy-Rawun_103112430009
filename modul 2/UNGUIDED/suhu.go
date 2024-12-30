package main

import "fmt"

func fahrenheitToCelsius(f float64) float64 {
  return (f - 32) * 5 / 9
}

func main() {
  var fahrenheit float64

  fmt.Print("Masukkan suhu dalam Fahrenheit: ")
  fmt.Scanln(&fahrenheit)

  celsius := fahrenheitToCelsius(fahrenheit)
  fmt.Printf("Suhu dalam Celsius: %.2f\n", celsius)
}