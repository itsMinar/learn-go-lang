package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang!")

	presentTime := time.Now()
	fmt.Println("Current Time: ", presentTime)

	fmt.Println("Formatted Time: ", presentTime.Format("01-02-2006 Monday 15:04:05"))

	createdDate := time.Date(1998, time.October, 13, 23, 23, 0, 0, time.UTC)
	fmt.Println("Created Date: ", createdDate.Format("01-02-2006 Monday"))
}
