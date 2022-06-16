package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your database host")
	var user *string = flag.String("user", "root", "Put your database user")
	var password *string = flag.String("password", "root", "Put your database password")
	var number *int = flag.Int("number", 100, "Put your number")

	flag.Parse() //untuk membca args -host , dll pada cli

	// flag berfungsi untuk membuat args agar rapih (biar otomatis ada minus2 nya)

	// fmt.Println("Host : ", host)
	// fmt.Println("User : ", user)
	// fmt.Println("Password : ", password)
	// fmt.Println("Number : ", number)

	// Host :  0xc00005e290
	// User :  0xc00005e2a0
	// Password :  0xc00005e2b0
	// Number :  0xc000012098

	//value ertrorkarena menggunakan pointer

	fmt.Println("Host : ", *host)
	fmt.Println("User : ", *user)
	fmt.Println("Password : ", *password)
	fmt.Println("Number : ", *number)

	// cara running :
	// go run .\50_udemy_flag.go -host="ekitesting.com" -user="ekitesting" -password="ekitesting123" -number=500

	// go run .\50_udemy_flag.go -host="ekitesting.com" -user="ekitesting"  (secara otomatis password,number menggungkana default)

	//https://pkg.go.dev/flag
}
