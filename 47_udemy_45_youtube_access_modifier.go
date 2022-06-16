package main //ini wajib
//package helper <<  ini bebas namanya

// 44. import

// D:\_GOPATH_GOLANG_PROJECT\src\_eki-latihan-golang-dasar\helper_44\sayhello.go
// D:\_GOPATH_GOLANG_PROJECT\src\_eki-latihan-golang-dasar\other_44\other.go

import (
	// "_eki-latihan-golang-dasar/helper_44"  // jika belum seting go modules  (file go.mod)  , _eki-latihan-golang-dasar adalah nama folder project yang ada pada folder src di gopath
	// "ekitesting/helper_44" // jika sudah seting go modules (file go.mod)

	// cara setting go modules
	// unset GOROOT
	// git init
	// go mod init ekitesting

	// go env -w GO111MODULE=off  // DISABLE GO MODULES
	// go env -w GO111MODULE=on  // ENABLE GO MODULES  <<< kalo go modules wajib pake ini

	"ekitesting/helper_44"
	"ekitesting/other_44"
	"fmt"
)

func main() {
	helper_44.SayHello("Eki")

	// helper.sayGoodbye("Eki") // error
	fmt.Println(helper_44.Application) // 45. access modifier - bs di akses dari luar (HURUF BESAR)
	other_44.SayHello("EKI")
	// helper_44.sayGoodbye("EKI") // 45. access modifier - sayGoodbye not exported by package helper_44compiler , akan error karena function di awalin huruf kecil
	// fmt.Println(helper.version) // error
}

// 45. access modifiler
// jika membuat function dengan huruf besar di awal artinya function dapat di panggil di semua package yang lain
// jika membuat function dengan huruf kecil di awal artinya function hanya dapat di panggil di package yang sama
