package main

import (
	"fmt"
	"strings"
)

func main() {

	// Beberapa Function di Package strings
	// Function Kegunaan
	// strings. Trim(string, cutset) Memotong cutset di awal dan akhir string
	// strings. ToLower(string) Membuat semua karakter string menjadi lower case
	// strings. ToUpper(string) Membuat semua karakter string menjadi upper case
	// strings.Split(string, separator) Memotong string berdasarkan separator
	// strings.Contains(string, search) Mengecek apakah string mengandung string lain
	// strings.ReplaceAll(string, from, to) = Mengubah semua string dari from ke to

	fmt.Println(strings.Contains("Indradi", "Eki"))
	fmt.Println(strings.Contains("Indradi", "Budi"))
	fmt.Println(strings.Split("Eki Indradi", " "))
	fmt.Println(strings.ToLower("Eki Indradi"))
	fmt.Println(strings.ToTitle("Eki Indradi"))
	fmt.Println(strings.ToUpper("Eki Indradi"))
	fmt.Println(strings.Trim("   Eki Indradi    ", " "))   // hapus spasi (kelebihan kiri  dan kanan aja)
	fmt.Println(strings.Trim("z   Eki Indradi    z", " ")) // hapus spasi (jika ada value sebelum kiri kanan spasi tidak akan hilang)
	fmt.Println(strings.ReplaceAll("eki aku eko eku eki aki", "eki", "EKI"))

}
