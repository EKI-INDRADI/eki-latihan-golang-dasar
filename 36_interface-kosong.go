package main

import "fmt"

// interface{} == interface kosong
// 1----------------------------------------
// func Ups() interface{} {

// 	return "Ups"
// }

// 1----------------------------------------

// 2----------------------------------------
func Ups(i int) interface{} { // kontrak yg tidak ada isi kontraknya , contoh bikin validasi data

	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}

}

// 2 ----------------------------------------

func main() {

	// contoh lain penggunaan
	// 1. fmt.Println(a...interface{})

	// 2. panic(y interface{})
	//    recover() interface{}

	// 1----------------------------------------
	// kosong := Ups()
	// fmt.Println(kosong)
	// 1----------------------------------------

	// 2----------------------------------------
	// var data int = Ups(1)  // gak bisa bgni (karena value berubah2 maka harus pake interface{})
	// harus value interface kosong , retrun nya jg harus interface kosong
	var data interface{} = Ups(1)
	fmt.Println(data)
	// 2----------------------------------------

}
