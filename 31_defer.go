package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	// pastikan selalu menggunakan defer  di atas
	defer logging() // pada process terakhir akan masuk ke logging function ( JIKA ERROR/NORMAL DI AKHIR SELALU DI EKSEKUSI )
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Result", result)
	// logging()  // JIKA TERJADI ERROR INI TIDAK AKAN PERNAH DI EKSEKUSI
}

func main() {
	runApplication(10) //OK
	//runApplication(0) //ERROR
}
