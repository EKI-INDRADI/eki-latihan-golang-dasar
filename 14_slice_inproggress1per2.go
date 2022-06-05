package main

import "fmt"

func main() {
	var months = [...]string{   // [...] <!---- DYNAMIC LENGHT

		// TEST months[4:7] 

		// slice low = skip data
		// slice high = mirip limit (low:hight)  <<<  low - height == LIMIT

		"Januari", // 0
		"Februari",// 1
		"Maret",// 2
		"April",// 3 
		"Mei",// 4   --> GET DATA  TEST months[4:7] 
		"Juni",// 5 --> GET DATA  TEST months[4:7] 
		"Juli",// 6 --> GET DATA  TEST months[4:7] 
		"Agustus",// 7
		"September",// 8
		"Oktober",// 9
		"November",// 10
		"Desember",// 11
	}

	//slice = mengakses data sebgian atau seluruh data di array 
	var slice1 = months[4:7]  // pointer 4 , length 3 (dari 4 ke 7), capacity 8
	//array[array index low:array index high]    = ambil data array index low sampe sebelum high
	//array[array index low:]  = ambil data array index low sampe akhir
	//array[:array index high]  = ambil data array index 0 sampe sebelum high
    //array[:]  = ambil dari index 0 sampe index sampe akhir


	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//months[5] = "Diubah"
	//fmt.Println(slice1)

	//slice1[0] = "Mei Lagi"
	//fmt.Println(months)

	var slice2 = months[11:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Eki")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Eki"
	newSlice[1] = "Indradi"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
