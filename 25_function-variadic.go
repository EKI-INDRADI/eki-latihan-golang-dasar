package main

// go: cannot find main module, but found .git/config in D:\_2021_test\belajar-golang-dasar
// 	to create a module there, run:
// 	go mod init
// go

import ( // auto import package jika pake go depedency visual studio code
	"fmt"
	"strconv"
)

// func sumAll(numbers []int) int { // bisa kyk gini
func sumAllx(numbers ...int) int { // bisa juga gini <===== VARIADIC/variable argument FUNCTION
	// cuma bisa 1 variable argument nya ,  hanya bisa dibuat di paling belakang (apabila ada lebih dr 1 argument variable)
	// cth :
	// func getFullName3(set_firstName string,set_middleName string,numbers ...int) (firstName string, middleName string, lastName string , total int) {
	// variable argument/variadic hanya bisa di paling belakang

	// function tipe data array (jika lebih dar 1 cukup gunakan tanda koma ",")
	// lebih tepatnya slice ( bukan array )
	total := 0
	for _, value := range numbers { // range numbers ini sama aja  []int
		// for _,      <=== tidak deklarasi variable maksudnya di awal maksudnya karena
		// sudah didekalarasi, total dan  range [10, 90, 30, 50, 40]  // mirip array.length     untuk kalkulasi datanya
		total += value
	}
	return total
}

func main() {
	// total := sumAll()  // variable argument ga harus ngisi
	total := sumAllx(10, 90, 30, 50, 40)
	fmt.Println(total)

	slice := []int{10, 90, 30, 50, 40} // jika ada tipe data slice / array
	total = sumAllx(slice...)          // cukup taro gini (memasukan slice data ke variadic function)

	fmt.Println(total)

	fmt.Println("slice data : ")
	fmt.Println(slice)

	fmt.Println("============================")
	fmt.Println("slice data variable argument : ")
	// fmt.Println(strings(slice...))

	console_vadiadic_value := make([]interface{}, len(slice))
	console_vadiadic_value2 := ""
	for i, v := range slice {
		console_vadiadic_value[i] = v
		console_vadiadic_value2 = console_vadiadic_value2 + strconv.Itoa(slice[i])

		if i < len(slice)-1 {
			console_vadiadic_value2 = console_vadiadic_value2 + ","
		}

		// fmt.Println(slice[i])

	}

	fmt.Println(console_vadiadic_value2)

	fmt.Println(console_vadiadic_value...)
}
