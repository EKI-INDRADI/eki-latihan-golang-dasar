package main

import "fmt"

func main() {

	// len(array)    		=== array length
	// array[index]  		=== array get index
	// array[index] = value === array set value

	var names [3]string

	names[0] = "Eki"
	names[1] = "Indradi"
	names[2] = "Khannedy"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string

	fmt.Println(len(lagi))
}
