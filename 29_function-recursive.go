package main

import "fmt"

//recusive function = function yang memanggil dirinya sendiri
// contoh pada kasus faktorial

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // result = result * i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1 // kalo revusive harus ada kondisi berhenti , kalo ga infinite loops / stack overflow
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	loop := factorialLoop(5)
	fmt.Println(loop) // ini pake looping
	// fmt.Println(5 * 4 * 3 * 2 * 1)

	recursive := factorialRecursive(5)
	fmt.Println(recursive) // ini pake recursif
}
