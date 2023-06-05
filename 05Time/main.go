package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime) //OP   2023-06-04 15:55:30.080352 +0530 IST m=+0.003616101

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //this date is a standart by Go

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC) //Date(year int, month time.Month,
	//day int, hour int, min int, sec int,
	//nsec int, loc *time.Location)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}

//go env
//go build
//GOOS="windows" go build
