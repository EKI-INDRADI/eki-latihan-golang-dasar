package main

import "fmt"

func random() interface{} {
	return "Eki"
}

// func main() {  //agar tidak terjadi panic saat mengunakaan type assetions ,
// 	result := random()
// 	switch value := result.(type) {   // contoh jika ingin ambil tipe data dari interface kosong
// 	case string:
// 		fmt.Println("String", value)
// 	case int:
// 		fmt.Println("Int", value)
// 	default:
// 		fmt.Println("Unknown")
// 	}
// }

func main() { // agar tidak terjadi panic saat mengunakaan type assetions ,
	// maka menggunakan switch untuk handle value (type assetions)

	//type assertions = merubah tipe data ke tipe data yang di inginkan (apabila betremu dengan interface kosong)
	var result interface{} = random()
	//var resultString string = result.(string)
	//fmt.Println(resultString)

	switch value := result.(type) { // contoh jika ingin ambil tipe data dari interface kosong
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}

//==================== PANIC ISSUE ====================

// func random() interface{} {
// //	return "OK"
// return true   <==== akan mengakibatkan panic
// }

// func main() {
// 	result := random()
// 	resultString := result.(string)
// 	fmt.Println(resultString)

// 	resultInt := result.(int) // panic
// 	fmt.Println(resultInt)
// }

//==================== /.PANIC ISSUE ====================
