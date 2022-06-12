package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) { // <=== ini adalah value asli yang akan berubah (country di main berubah jadi indonesia)
	//func ChangeCountryToIndonesia(address Address) { <=== ini adalah value duplicate nya yg berubah (countery tidak ada)
	address.Country = "Indonesia"
}

func main_pass_by_value() {
	// di golang variable (default) = pass by value  (copy value ke variahlle baru)
	// artinya copy memrory / duplicate data, jika ada perubahan pada salah 1 variable ,
	// maka tidak akan mempengaruhi variable lain yang di copy
	// case ini mirip javascript (JSON.parse(JSON.stringify(obj_or_array_value)))

	addressl := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := addressl

	address2.City = "Bandung"

	fmt.Println(addressl) // address1 tidak berubah
	fmt.Println(address2)

	// jika pada javascript (variable biasa)
	// let addressl = ["Subang", "Jawa Barat", "Indonesia"]
	// let address2 = addressl
	// address2.push("Bandung")
	// jika address2 berubah maka address1 juga berubah
	// resultnya addressl = ["Subang", "Jawa Barat", "Indonesia", "Bandung"]
	// resultnya address2 = ["Subang", "Jawa Barat", "Indonesia", "Bandung"]

	// jika di golang mirip javascript (JSON.parse(JSON.stringify(obj_or_array_value)))
	// addressl := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := addressl
	// address2.City = "Bandung"
	// resultnya addressl = {"Subang", "Jawa Barat", "Indonesia"}
	// resultnya address2 = {"Bandung", "Jawa Barat", "Indonesia"}

	// result
	// {Subang Jawa Barat Indonesia}
	// {Bandung Jawa Barat Indonesia}

}

func main_pass_by_reference() { // bukan duplicate / replace value / di memory yang sama
	addressl := Address{"Subang", "Jawa Barat", "Indonesia"}

	address2 := &addressl // & = POINTER =  membuat pointer agar value pass by reference / replace value / memory data sama

	// address2 := &addressl    sama dengan     address2 *address2 = &addressl   (menggunakan tipe data)   ( & = POINTER , * = POINTER)

	// ini kalo di masukan tipe datanya sama saja

	// var addressl Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	// var address2 *Address= &addressl

	// Operator * = penanda itu adalah pointer ke data yang mana

	address2.City = "Bandung" // jika address2 berubah maka address1 juga berubah

	fmt.Println(addressl) // address1 valuenya akan sama dengan address2
	fmt.Println(address2)

	// case ini mirip default variablenya javascript / typescript

	// 	result :
	// {Bandung Jawa Barat Indonesia}
	// &{Bandung Jawa Barat Indonesia} // agar tidak muncul & di depan

}

func main() { //
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	// var address4 *Address = &Address{"Subang", "Jawa Barat", "Indonesia"}  // BIKIN DATA BARU (YG BERISI DATA)
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Bandung"

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"} // jika begini semua address1 (memory) di paksa pindah ke address2 (memory) , artinya address3 pun mengikuti address2
	// itulah fungsi dari *   (mengacu ke memory yg paling baru)

	// 11:00 https://www.youtube.com/watch?v=uhF5MVmYwA4&list=PL-CtdCApEFH_t5_dtCQZgWJqWF45WRgZw&index=40

	fmt.Println(address1) //{"Malang", "Jawa Timur", "Indonesia"}
	fmt.Println(address2) //&{"Malang", "Jawa Timur", "Indonesia"}
	fmt.Println(address3) //&{"Malang", "Jawa Timur", "Indonesia"}

	var address4 *Address = new(Address) // membuat pointer baru , bedanya new hanya mengembalikan pointer ke data kosong (tidak ada awal data)
	address4.City = "Jakarta"
	fmt.Println(address4)

	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}
	// ChangeCountryToIndonesia(alamat) // bakalan error karena alamat bukan pointer
	// 	ChangeCountryToIndonesia(&lamat) //tidak error karena menggunakan pointer
	var alamatPointer *Address = &alamat
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)
}
