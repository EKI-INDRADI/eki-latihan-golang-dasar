package main

import "fmt"

//-----------------------------------1 PANIC

func endApp() {
	//-----------------------------------2 RECOVER
	message := recover() // untuk menangkap erornya (messsage) bisa tangkap pake recover
	// tanpa recocver tidak bs tangkap error
	//  dan APLIKASI AKAN TETAP BERJALAN tidak terjadi aplikasi berhenti ( mirip try catch JS/TS)
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	//-----------------------------------2 RECOVER
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
		// function ini memaksa program berhenti
		// ( kaya throw new error di javascript ), dan defer function akan selalu di eksekusi
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Eki")
}

//-----------------------------------1 PANIC
