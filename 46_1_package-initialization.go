package main

import (
	"ekitesting/database_46" //jika import tidak di gunakan,  import akan hilang
	"fmt"
)

// import (
// 	_ "ekitesting/database_46" // 46. blank identifier , jika import tidak di gunakan tidak akan hilang
// 	"fmt"
// )

func main() {
	result := database_46.GetDatabase()
	fmt.Println(result)
}

//46. package initialization
// func init()  ==== untuk membuat yang  di akses secara otomatis ketika pakage di akses, kita cukup membuat function dengan nama init
// D:\_GOPATH_GOLANG_PROJECT\src\_eki-latihan-golang-dasar\database_46\database_46.go
