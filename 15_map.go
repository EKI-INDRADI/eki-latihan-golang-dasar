package main

import "fmt"

func main() {

	// var person = map[string]string{
	// 	"name":    "Eki",
	// 	"address": "Subang",
	// }

	// map mirip array , atau lebih mirip object

	person := map[string]string{
		"name":    "Eki",
		"address": "Subang",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Eki"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book)) 
	// len(map) // mengambil jumlah data pada map nya
	// map[key] // mengambil value dari key nya
	// delete(map, key) // menghapus data map dgn key
	// map[key] = value // mengisi value berdasarkan key
	// make[map[TypeKey]TypeValue] // membuat map baru


	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
