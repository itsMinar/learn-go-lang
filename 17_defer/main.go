package main

import "fmt"

func main() {
	defer fmt.Println("This is 1st Defer statement.")
	defer fmt.Println("This is 2nd Defer statement.")
	defer fmt.Println("This is 3rd Defer statement.")

	fmt.Println("Welcome to Defer in Go!")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
