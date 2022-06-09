package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {

	//  Pembagi(nilai int, pembagi int) == value input
	// (int, error)  = return balikan
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0") // return valikan sesuai tipe data
	} else {
		result := nilai / pembagi
		return result, nil // return valikan sesuai tipe data , karena error berupa interface maka bs gunakan nil (undefined/null di golang)
	}
}

func main() {

	hasil, err := Pembagi(100, 0) // 	hasil, err  = return balikan , golang bisa 2 value balikan
	if err == nil {               // jika errornya == nil
		fmt.Println("Hasil", hasil)
	} else { // jika errornya != nil   / ada value error
		fmt.Println("Error", err.Error())
	}
}
