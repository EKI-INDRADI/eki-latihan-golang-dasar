package main

import "fmt"

func main() { // closure = berinteraksi dengan data yang ada di sekitarnya   dari { sampe }
	name := "Eki"
	counter := 0 // 2. dengan ini  , ini akan menjadi closure ( di baca di dalam function)

	increment := func() {
		// counter := 1  // supaya gak bug harus bikin ini agar gak jadi closure (agar yidak bca dari var luar)
		name := "Budi"
		fmt.Println("Increment")
		counter++ // 1. ini harus hati2 dengan no 2
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
