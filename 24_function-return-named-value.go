package main // dalam 1 package main tidak boleh buat nama function yang sama akan error

import "fmt"

func getFullName3() (firstName string, middleName string, lastName string) {
	firstName = "Eki"
	middleName = "Indradi"
	lastName = "test3"

	// return firstName,middleName,lastName // return secara explisit
	return // cukup return saja , agar bisa rubah variablenya (sesuai urutan variable deklarasi =  (firstName string, middleName string, lastName string))

	// ini cuma ada di golang
}

func main() {

	// firstName,middleName,lastName := getFullName2()
	// fmt.Println(firstName,middleName,lastName)

	a, b, c := getFullName3() // bisa rubah isi variablenya
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
