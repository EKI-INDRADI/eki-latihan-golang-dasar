package main

import "fmt"

type Man struct {
	Name string
}

// func (man Man) Married() { // tanpa pointer
// 	man.Name = "Mr. " + man.Name
// }

func (man *Man) Married() { //dengan pointer , kelebihan lebih hemat memory
	man.Name = "Mr. " + man.Name
}

func main() {
	Eki := Man{"Eki"}
	Eki.Married()

	fmt.Println(Eki.Name)
}
