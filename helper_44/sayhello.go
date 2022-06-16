package helper_44

import "fmt"

var version = 1                    // 45. access modifier - tidak bisa di akses dari luar
var Application = "Belajar Golang" // 45. access modifier -  bs di akses dari luar

func SayHello(name string) { // NAMA FUNCTION TIDAK BOLEH SAMA, DALAM 1 FOLDER AKAN ERROR (JIKA PACKAGE SAMA)
	// 47. access modifier - bs di akses dari luar

	fmt.Println("Hello", name)
}

func sayGoodbye(name string) { // NAMA FUNCTION TIDAK BOLEH SAMA, DALAM 1 FOLDER AKAN ERROR (JIKA PACKAGE SAMA)
	// 47. access modifier -  tidak bisa di akses dari luar
	fmt.Println("Goodbye", name)
}
