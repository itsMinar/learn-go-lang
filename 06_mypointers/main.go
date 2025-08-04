package main

import "fmt"

func main() {
	fmt.Println("Welcome to a Class on Pointers in Go!")

	// var ptr *int
	// fmt.Println("Value of pointer is:", ptr)

	myNumber := 23
	var ptr *int = &myNumber
	fmt.Println("Value of pointer is:", ptr)
	fmt.Println("Value of myNumber is:", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value of myNumber is:", myNumber)
	fmt.Println("New value of pointer is:", *ptr)
}
