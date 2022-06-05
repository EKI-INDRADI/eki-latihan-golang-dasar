package main

import "fmt"


func getFullName() (string, string, string) { 
	return "Eki", "Indradi", "test3"
}

func getFullName2() (string, string, string, string , int) { 
	// func golang bisa handle multiple value  (cuma ada di golang)
	// (string, string, string, string , int) {   <<< RETURN VALUE TIPE DATA WAJIB DEKLARASI
	
	
	return "Eki", "Indradi", "test3", "test4" , 5
}

func main() {
	firstName, _, _ := getFullName()   // _ untuk gak peduli sama value return nya  (jadi gak wajib ngambil)
	// firstName, _, _ := getFullName() // minimal harus ada 1 value return , kalo _, semua bakalan error
	fmt.Println(firstName)

	ekitesting1, ekitesting2, ekitesting3, ekitesting4, ekitesting5 := getFullName2()
	fmt.Println(ekitesting1)
	fmt.Println(ekitesting2)
	fmt.Println(ekitesting3)
	// fmt.Println(ekitesting4) // kalo gak pake _ , maka ekitesting4 declared but not used ( harus di tangkep kalo ga pake _ )
	fmt.Println(ekitesting4)
	fmt.Println(ekitesting5)


	// fmt.Println(middleName)
	// fmt.Println(lastName)

	//  go run .\23_function-return-multiple-value.go -o .\_GO_BUILD
}
