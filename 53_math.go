package main

import (
	"fmt"
	"math"
)

func main() {

	//https://pkg.go.dev/math
	// math.Round(float64) Membulatkan floaté4 keatas atau kebawah, sesuai. dengan yang paling dekat
	// math.Floor(float64) Membulatkan floaté4 kebawah
	// math.Ceil(float64) Membulatkan floaté4 keatas
	// math.Max(float64, float64) Mengembalikan nilai float64 paling besar
	// math.Min(float64, float64) Mengembalikan nilai floaté4 paling kecil

	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.3))
	fmt.Println(math.Floor(1.7))
	fmt.Println(math.Ceil(1.3))

	fmt.Println(math.Max(10, 20))
	fmt.Println(math.Min(10, 20))
}
