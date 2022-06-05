package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

//================== tidak perlu deklarasi jika menggunakan anonymous function
//func blacklistAdmin(name string) bool {
//	return name == "admin"
//}
//
//func blacklistRoot(name string) bool {
//	return name == "root"
//}
//================== tidak perlu dekalrasi jika menggunakan anonymous function

func main() {

	// anonymous function / bikin function secara langsung , tanpa hrs buat function terlebih dahulu

	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("Eki", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("Eki", func(name string) bool {
		return name == "root"
	})
}
