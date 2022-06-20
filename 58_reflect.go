package main

import (
	"fmt"
	"reflect"
)

// refect = saat compile/runtime kita dapat membaca struktur code

type Sample struct {
	Name string `required:"true" max:"10"`
}

type ContohLagi struct {
	Name        string `required:"true"`
	Description string `required:"true"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	// https://pkg.go.dev/reflect
	sample := Sample{"Eki"}

	var sampleType reflect.Type = reflect.TypeOf(sample)

	// reflaction ini berguna untuk validasi data , jadi ga perlu bikin if banyak2
	// contoh validasi func IsValid

	fmt.Println("sampleType.NumField() ", sampleType.NumField())
	fmt.Println("sampleType.Field(0).Name ", sampleType.Field(0).Name) //Field(0) = Name // type Sample struct {
	fmt.Println("sampleType.Field(0).Tag.Get('required')) ", sampleType.Field(0).Tag.Get("required"))
	fmt.Println("sampleType.Field(0).Tag.Get('max') ", sampleType.Field(0).Tag.Get("max"))

	sample.Name = ""
	fmt.Println("IsValid(sample) ", IsValid(sample))

	sample.Name = "eki"
	fmt.Println("IsValid(sample) eki : ", IsValid(sample))

	contoh := ContohLagi{"Eki", "Oke"}
	fmt.Println("IsValid(contoh) ", IsValid(contoh))
}
