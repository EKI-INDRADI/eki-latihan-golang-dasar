package main

import "fmt"

func NewMap(name string) map[string]string { // default value bisa di set kalo di  golang , kalo di js kan undefined / null
	// nil = null

	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {

	//===================== tanpa nil =====================

	// var person map[string]string

	// if person["name"] == "" {
	// 	fmt.Println("Data Kosong")
	// } else {
	// 	fmt.Println(person)
	// }

	//===================== tanpa nil =====================

	//===================== func nil =====================

	// nil hanya bisa di gunakan pada fungsi interface, function , map , slice , pointer dan channel

	var person map[string]string = NewMap("Eki")

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}

	//===================== func nil =====================
}
