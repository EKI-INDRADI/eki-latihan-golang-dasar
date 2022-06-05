package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Eki Indradi", "Eki"))
	fmt.Println(strings.Contains("Eki Indradi", "Budi"))

	fmt.Println(strings.Split("Eki Kurniawna Khannedy", " "))

	fmt.Println(strings.ToLower("Eki Indradi Khannedy"))
	fmt.Println(strings.ToUpper("Eki Indradi Khannedy"))
	fmt.Println(strings.ToTitle("Eki Indradi khannedy"))

	fmt.Println(strings.Trim("      Eki Indradi     ", " "))
	fmt.Println(strings.ReplaceAll("Eki Joko Eki", "Eki", "Budi"))
}
