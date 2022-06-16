package main

import (
	"fmt"
	"strconv"
)

func main() {
	//strconv == string convertion

	//https://pkg.go.dev/strconv

	// Beberapa Function di Package strconv
	// strconv.parseBool(string) Mengubah string ke bool
	// strconv.parseFloat(string) Mengubah string ke float
	// strconv.parselnt(string) Mengubah string ke int64
	// strconv.FormatBool(bool)
	// strconv.FormatFloat(float, ... ) Mengubah float64 ke string
	// strconv.Formatint(int, ... ) Mengubah int64 ke string

	boolean, err := strconv.ParseBool("true")
	// hampir setiap strconv mengembalikan 2 value (data,error)

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	//========================= ERROR
	// boolean1, err1 := strconv.ParseBool("salah")
	// hampir setiap strconv mengembalikan 2 value (data,error)

	// 	if err1 == nil {
	// 		fmt.Println(boolean1)
	// 	} else {
	// 		fmt.Println(err.Error())
	// 	}

	// 	panic: runtime error: invalid memory address or nil pointer dereference
	// [signal 0xc0000005 code=0x0 addr=0x18 pc=0x40c888]

	// goroutine 1 [running]:
	//========================= ERROR

	number, err := strconv.ParseInt("1000000", 10, 64)
	// number, err := strconv.ParseInt("salah", 10, 64) //strconv.ParseInt: parsing "salah": invalid syntax
	//("x,VALUE,x) == 10 = decimal , 2 = binary , 8 = octal ,16 = hexadecimal
	//("x,x,VALUE) == 64 = int 64 , 32 = int 32 , 16 = int 16 , 8 = int 8
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 10) //10 = decimal , 2 = binary , 8 = octal ,16 = hexadecimal
	fmt.Println(value)

	valueInt, _ := strconv.Atoi("2000000") // == strconv.ParseInt("1000000", 10, 0)
	fmt.Println(valueInt)
}
