package main

import "fmt"

func getHello(name string) string {     // (..) string {   <==== return harus tipe data string
	if name == "" { 
		return "Hello Bro"
	} else {
		return "Hello " + name
	}

	//setelah return data tidak akan di eksekusi ( mirip  cara return javascript/typescript)
}

func main() {
	result := getHello("Eki")
	fmt.Println(result) //console.log hasil return
	fmt.Println(getHello("")) //console.log hasil return
}
