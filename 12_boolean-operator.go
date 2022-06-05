package main

import "fmt"

func main() {

	// && dan  (pebndingan 2 bool)
	// || atau (pebndingan 2 bool)
	// ! kebalikan boolean (hanya untuk 1 boolean)

	// true && true      =    true
	// true && false     =    false
	// false && true     =    false
	// false && false    =    false

	// true || true    = true
	// true || false   = true
	// false || true   = true
	// false || false  = false

	// !  true  = false
	// !  false = true


	

	var ujian = 80
	var absensi = 75

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusAbsensi && lulusUjian
	fmt.Println(lulus)

	fmt.Println(ujian >= 80 && absensi >= 80)

}
