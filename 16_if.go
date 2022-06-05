package main

import "fmt"

func main() {
	var name = "Indradi"

	if name == "Eki" {
		fmt.Println("Hello Eki")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hi, kenalan donk")
	}


	// var length = len(name)
	// if length > 5 {  // tanpa short statement
	// 	fmt.Println("Terlalu Panjang")
	// } else {
	// 	fmt.Println("Nama sudah benar")
	// }

	if length := len(name); length > 5 //short statement di eksekusi sebelum melakukan if
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}



}
