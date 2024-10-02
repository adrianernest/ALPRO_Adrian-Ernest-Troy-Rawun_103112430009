package main

import (
	"fmt"
	"math"
)

func main() {
	var jarijari int

	fmt.Print("Jejari = ")
	fmt.Scanln(&jarijari)

	volume := (4.0 / 3.0) * math.Pi * math.Pow(float64(jarijari), 3)
	luasPermukaan := 4 * math.Pi * math.Pow(float64(jarijari), 2)

	fmt.Printf("Bola dengan jejari %d memiliki volume %.3f dan luas kulit %.3f\n", jarijari, volume, luasPermukaan)
}