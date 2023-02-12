package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to TimeStudy in goLang ")

	PresentTime := time.Now() //time module is used here for time
	fmt.Println(PresentTime)

	fmt.Println(PresentTime.Format("01-02-2006 15:04:05 Monday")) //this should be written like this is you want the correct output

	currentDate := time.Date(2023, time.August, 15, 23, 59, 59, 99, time.UTC)
	fmt.Println(currentDate)
	fmt.Println(currentDate.Format("01-02-2006 Monday"))

}
