package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	// regular expression
	// untuk menggunakan pencarian

	// https://github.com/google/re2/wiki/Syntax // cara buat syntax regex
	// https://pkg.go.dev/regexp

	// Beberapa Function di Package regexp
	// regexp.MustCompile(string) membuat regexp
	// Regexp.MatchString(string) bool Mengecek apakah Regexp match dengan string
	// Regexp.FindAllString(string, max) | Mencari string yang match dengan maximum jumlah hasil

	fmt.Println(regex.MatchString("Eki"))
	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("eDo"))

	// var regex2 = regexp.MustCompile(`e([a-z])o`)
	// fmt.Println(regex2.MatchString("eko"))
	// fmt.Println(regex2.MatchString("edo"))
	// fmt.Println(regex2.MatchString("eKo"))

	search := regex.FindAllString("Eki eka edo eto eyo eki", -1) //-1 semuanya

	// search := regex.FindAllString("Eki eka edo eto eyo eki", 2)
	// search := regex.FindAllString("Eki eka edo eto eyo eki", 1)
	fmt.Println(search)
}
