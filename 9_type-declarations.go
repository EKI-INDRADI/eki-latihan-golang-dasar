package main

import "fmt"

func main() { 
	type NoKTP string  
	//  <!---- tipe data baru
	type Married bool



	// bikin tipe data alias yang sudah ada

	var noKtpEki NoKTP = "18741982741897419874"
	var marriedStatus Married = true
	fmt.Println(noKtpEki)
	fmt.Println(marriedStatus)
}
