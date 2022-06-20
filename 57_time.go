package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println("now", now)
	fmt.Println("now.Year()", now.Year())
	fmt.Println("now.Month()", now.Month())
	fmt.Println("now.Day()", now.Day())
	fmt.Println("now.Hour()", now.Hour())
	fmt.Println("now.Minute()", now.Minute())
	fmt.Println("now.Second()", now.Second())
	fmt.Println("now.Nanosecond()", now.Nanosecond())

	// SEMUA FORMAT DI GOLANG HARUS PAKE TANGGAL LAGI , BUKAN yyyy-MM-dd HH:mm:ss

	//Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location)
	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println("utc", utc)

	layout := "2006-01-02" //"yyyy-MM-dd"  cth kalo di javascript/ts
	//layout := time.RFC3339 // (ini sama aja kayak toISOString() di javascript )
	// 	RFC3339     = "2006-01-02T15:04:05Z07:00"
	//  layout := "2006-01-02"  //2006-01-02 <<< default dari RFC itu 2006-01-02 ,
	// jadi kalo input ini sama aja ngikuti standar parse RFC

	// "2006-01-02"  == 1990-03-20 00:00:00 +0000 UTC

	parse, _ := time.Parse(layout, "1990-03-20") //C:\Program Files\Go\src\time\format.go
	fmt.Println("parse ", parse)

	// ------------------ WINDOWS WITHOUT ANTIVIRUS ISSUE
	// 	go build -o .\_GO_BUILD\ 57_time.go
	//  start .\_GO_BUILD\57_time.exe
	// ------------------ WINDOWS WITHOUT ANTIVIRUS ISSUE

	// format  ISO 8601 // toISOString() javascript
	// fmt.Println("ISO 8601", time.Now().Format(time.RFC3339))

	fmt.Println("ISO 8601 (javascript .toISOString() ): ", time.Now().UTC().Format("2006-01-02T15:04:05-0700"))
}
