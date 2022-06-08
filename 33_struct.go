package main

import "fmt"

type Customer struct {
	// mirip class fungsinya
	// template data / prototype data
	// mirip interface di typescript / class di typescript
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuuuuu from", a.Name)
}

func main() {
	var Eki Customer
	Eki.Name = "Eki"
	Eki.Address = "Indonesia"
	Eki.Age = 30

	Eki.sayHi("Joko")
	Eki.sayHuuu()

	//fmt.Println(Eki.Name)
	//fmt.Println(Eki.Address)
	//fmt.Println(Eki.Age)
	//
	//joko := Customer{
	//	Name:    "Joko",
	//	Address: "Cirebon",
	//	Age:     35,
	//}
	//fmt.Println(joko)
	//
	//budi := Customer{"Budi", "Jakarta", 35}
	//fmt.Println(budi)
}
