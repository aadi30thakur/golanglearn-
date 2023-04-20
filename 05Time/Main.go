package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the time study")
	presentTime := time.Now()
	fmt.Println("The time is", presentTime)
	fmt.Println("The year is", presentTime.Format("01-02-2005Z-10:04:05PMT"))

	createddate := time.Date(2020, time.January, 12, 15, 34, 50, 0, time.UTC)
	fmt.Println("The created date is", createddate)
	//formated date
	fmt.Println("The formated date is", createddate.Format("01-02-2006 Monday"))
}
