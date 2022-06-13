package database_46

import "fmt"

var connection string

func init() { //46. package initialization
	fmt.Println("Init dipanggil")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
