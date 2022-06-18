package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

// ------------------ FUNCTION INI DI BUAT BERDASARKAN INTERFACE SORT  / SORT FUNCTION
func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {

	// untuk balikan data bisa dibuat bikin temporary data , tetapi  ada cara yang lebih cepat :

	value[i], value[j] = value[j], value[i]

	// jadi seakan akan
	// value[i] akan di ambil darri value [j]
	// value[j] akan di ambil darri value [i]

}

// ------------------ /FUNCTION INI DI BUAT BERDASARKAN INTERFACE SORT  / SORT FUNCTION

// function di buat untuk memudahkan sort data

func main() {
	// https://pkg.go.dev/sort
	// sort data

	// INTERFACE SORT : C:\Program Files\Go\src\sort\sort.go
	// 	// elements of the collection be enumerated by an integer index.
	//  type Interface interface {
	// 	// Len is the number of elements in the collection.
	// 	Len() int    // jumlah data
	// 	// Less reports whether the element with
	// 	// index i should sort before the element with index j.
	// 	Less(i, j int) bool // apakah i lebih kecil dari J jika benar ==  true
	// 	// Swap swaps the elements with indexes i and j.
	// 	Swap(i, j int) // swap data dengan index i dan j
	// 	}

	users := []User{
		{"Eki", 27},
		{"Jokok", 35},
		{"Budi", 31},
		{"Rudi", 25},
	}

	fmt.Println(users)
	// sort.Sort(users) // ini akan error karena function sort butuh data interface, karena []User tidak punya struct / contract / interface
	sort.Sort(UserSlice(users)) //solusinya pake alias  //sort.Sort(UserSlice(users))  , karna UserSlice ngikutin contact interface
	fmt.Println(users)
}
