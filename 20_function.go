package main

import "fmt"

func sayHello() { 
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
}

func main() { // function default  (mirip form_load di vb net / gOnInit() di angular) , 
	//golang megunakan case sensitive pattern
	for i := 0; i < 10; i++ {
		sayHello()
	}
}
