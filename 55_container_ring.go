package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {

	// ------------------ WINDOWS WITHOUT ANTIVIRUS ISSUE
	// 	go build -o .\_GO_BUILD\ 54_container-list.go
	//  start .\_GO_BUILD\54_container-list.exe
	// ------------------ WINDOWS WITHOUT ANTIVIRUS ISSUE

	// packge container ring = struktur data circular list
	// setiap next/previusnya (jika habis data akan kembali ke awal, selalu berulang)

	var data *ring.Ring = ring.New(5) // data wajib di isi lenght nya

	//========== contoh masukin data
	// data.Value = "eki"

	// // var data2 = data.Next()
	// data2.Value = "indradi"
	//========== /contoh masukin data

	for i := 0; i < data.Len(); i++ { //memasukan data by loop
		data.Value = "Data " + strconv.FormatInt(int64(i), 10)
		data = data.Next() //biar ke ring selanjutnya
	}

	// fmt.Println(*data)

	data.Do(func(value interface{}) { // cara print value pake func do bawaan ring
		fmt.Println(value) // function ring  << ada do untuk print semua element yang ada pada ring tsb
	})

	// https://pkg.go.dev/container/ring
}
