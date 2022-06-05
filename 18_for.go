package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		// init statement = counter := 1  (eksekusi di awal/initialisasi), 
		// condition =  counter <= 10
		// post statement = counter++  (eksekusi di akhir)
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"Eki", "Indradi", "Khannedy", "Budi", "Joko"}


	// for counter <= 10 {
	// fmt.Println("Perulangan ke ", counter)
    //    counter++
	// }

	for i := 0; i < len(slice); i++ { // perulangan di goolang hanya ada 1, for loops (tidak ada do while , while, dll)
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Eki"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
