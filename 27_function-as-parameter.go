package main

import (
	"fmt"
	"strings"
)

//===================v1====================

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	// sebelum print maka akan filter dulu menggunakan function yang ada pada
// 	//  func sayHelloWithFilter(............... filter func(string) string) {
// 	fmt.Println("Hello", nameFiltered)
// }

//===================v1====================

//===================v2====================

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

//===================v2====================

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func upperFilter(name string) string {
	return strings.ToUpper(name)
}

func sayHelloWithFilter3(name string, filter func(string) string, filter2 func(string) string) {
	nameFiltered := filter(name)
	nameFiltered2 := filter2(nameFiltered)
	// sebelum print maka akan filter dulu menggunakan function yang ada pada
	//  func sayHelloWithFilter(............... filter func(string) string) {
	fmt.Println("nameFiltered -> nameFiltered2" + " | " + nameFiltered + " -> " + nameFiltered2)
}

func main() {
	sayHelloWithFilter("Eki", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
	sayHelloWithFilter3("eki", spamFilter, upperFilter)
	sayHelloWithFilter3("Anjing", spamFilter, upperFilter)
	// tujuan bikin function as parameter adalah supaya gak terlalu bnyk if bersarang
}
