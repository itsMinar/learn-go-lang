package main

// code segment - all functions
// data segment - global memory
// stack
// heap

import "fmt"

var a = 10

func add(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println("Sum is:", sum)
}

func main() {
	fmt.Println("Internal Memory in Go lang")

	add(2, 5)
}

func init() {
	fmt.Println("Welcome")
}
