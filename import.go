package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	helper.SayHello("Eki")
	// helper.sayGoodbye("Eki") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error
}
