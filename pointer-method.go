package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	Eki := Man{"Eki"}
	Eki.Married()

	fmt.Println(Eki.Name)
}
