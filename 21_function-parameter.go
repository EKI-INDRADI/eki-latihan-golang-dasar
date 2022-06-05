package main

import "fmt"

func sayHelloTo(firstName string, lastName string) { // parameter
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "Eki"
	sayHelloTo(firstName, "Indradi")
	sayHelloTo("ek2", "indradi2")
}
