package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Functions in Go!")
	greeter("Hello, World!")

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult := proAdder(10, 20, 30, 40, 50)
	fmt.Println("Pro Result is: ", proResult)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func greeter(msg string) {
	fmt.Println("Hello,", msg)
}
