package main

import "fmt"

func main() {

	name := "Indradi"

	switch name {
	case "Eki":
		fmt.Println("Hello Eki")
		fmt.Println("Hello Eki")
	case "Joko":
		fmt.Println("Hello Joko")
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Kenalan Donk")
		fmt.Println("Kenalan Donk")
	}

	//switch length := len(name); length > 5 { // sort statement di eksekusi sebelum melakukan if
	//case true:
	//	fmt.Println("Nama terlalu panjang")
	//case false:
	//	fmt.Println("Nama sudah benar")
	//}

	length := len(name) // tanpa sort statement 
	switch { // di golang bs tanpa switch codition apapun
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
