package main

import (
	"container/list"
	"fmt"
)

func main() {

	//packge container.list = implementasi struktur data double linked list
	//data rata yang saling keterkaitan setelah atau sebelumnya ,
	//mirip seperti array /slice tetapi bedanya ga punya index
	// jadi urutan data akan saling keterkaitan ,
	// di awal selalu kosong
	// jadi  nil , 12,9,1231,123,123,nil
	// urutan nya saling berkaitan dan urut (data sebelum dan setelahnya)

	// double linked list = ada keterkaitan dengan prev dan next  dari value yang di input

	// data awal  selalu nil
	// data akhir selalu nil

	// linked list tidak ada index kyk array

	// https://pkg.go.dev/container/list

	data := list.New() // bikin list baru

	data.PushBack("Eki")
	// PushBack = menambahkan data ke next nya (ibaratnya kalo array bakalan langsung di array index terakhir)
	// data di input dari belakang
	data.PushBack("Indradi")
	data.PushBack("xxx")
	data.PushFront("Budi")
	// PushFront = data di input dari depan ( mirip array index 0 / array.push di js)

	// data.Front().Value // ambil data paling awal
	fmt.Println("==================data.Front().Value : ==================")
	fmt.Println(data.Front().Value)
	fmt.Println("==================")

	fmt.Println("==================data.Front().Next().Value : ==================")
	fmt.Println(data.Front().Next().Value)
	fmt.Println("==================")

	fmt.Println("==================data.Front().Next().Next().Next().Value : ==================")
	fmt.Println(data.Front().Next().Next().Next().Value)
	fmt.Println("==================")

	fmt.Println("==================data.Front().Prev() (nil) : ==================")
	fmt.Println(data.Front().Prev()) // selalu nill di awal
	fmt.Println("==================")

	fmt.Println("==================data.Front().Next() (nil)  : ==================")
	fmt.Println(data.Front().Next())
	fmt.Println("==================")

	// ------------------ WINDOWS WITHOUT ANTIVIRUS ISSUE
	// 	go build -o .\_GO_BUILD\ 54_container-list.go
	//  start .\_GO_BUILD\54_container-list.exe
	// ------------------ WINDOWS WITHOUT ANTIVIRUS ISSUE

	// dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
