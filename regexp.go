package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("Eki"))
	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("eDo"))

	search := regex.FindAllString("Eki eka edo eto eyo eki", -1)
	fmt.Println(search)
}
