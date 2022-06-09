package main

import "fmt"

type Customer struct {
	// mirip class fungsinya
	// template data / prototype data
	// mirip interface di typescript / class di typescript
	Name, Address string
	Age           int
}

//======================================== STRUCT METHOD =========================================

// mirip membuat function di interface typescript
func (customer Customer) sayHi(name string) { // cara panggil struct di dalam function (membuat alias)
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHuuu() { // cara panggil struct di dalam function (membuat alias)
	fmt.Println("Huuuuuu from", a.Name)
}

//======================================== /STRUCT METHOD =========================================

func main() {
	var Eki Customer // (eki alias customer)
	Eki.Name = "Eki"
	Eki.Address = "Indonesia"
	Eki.Age = 30

	//======================================== STRUCT METHOD =========================================

	Eki.sayHi("Joko") // (seakan akan memanggil function sayHi di dalam struct customer)
	Eki.sayHuuu()     //  (seakan akan memanggil function sayHuuu di dalam struct customer)

	//======================================== /STRUCT METHOD =========================================

	//fmt.Println(Eki.Name)
	//fmt.Println(Eki.Address)
	//fmt.Println(Eki.Age)

	//  contoh lain :
	//joko := Customer{
	//	Name:    "Joko",
	//	Address: "Cirebon",
	//	Age:     35,
	//}
	//fmt.Println(joko)
	//
	//budi := Customer{"Budi", "Jakarta", 35} // harus urut datanya sesuai dengan structnya , dan harus mengisi seluruh struct value
	//fmt.Println(budi)
}
