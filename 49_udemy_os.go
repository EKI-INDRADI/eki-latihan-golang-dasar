package main

import (
	"fmt"
	"os"
)

func main() {
	// pakcage OS untuk mengoperasikan sistem operasi

	args := os.Args
	fmt.Println("Argument : ")
	fmt.Println(args)

	fmt.Println("LOKASI EXECUTE FILE DENGAN ARGUMENT :", args)

	// contoh
	// go run .\47_os.go eki tamfan
	// args itu sama aja "eki tamfan"

	fmt.Println("ARGUMENT[0] :", args[0])

	fmt.Println("ARGUMENT[1] :", args[1])

	fmt.Println("ARGUMENT[2] :", args[2])

	// fmt.Println("LOKASI EXECUTE FILE DENGAN ARGUMENT[3] :", args[3]) ini akan error

	name, err := os.Hostname() // pda function Hostname ada 2 balikan, 1 string dan 2 error
	if err == nil {            // jika error == undefined / null
		fmt.Println("Hostname : ", name)
	} else {
		fmt.Println("Error", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	// example untuk setup ENV di dalam OS untuk linux dan mac
	// export APP_USERNAME=eki
	// export APP_PASSWORD=tamfan

	//for windows https://stackoverflow.com/questions/559816/how-to-export-and-import-environment-variables-in-windows
	// C:\> SET >> allvariables.txt
	// C:\> for /F %A in (allvariables.txt) do SET %A

	// https://pkg.go.dev/os
	fmt.Println(username)
	fmt.Println(password)
}
