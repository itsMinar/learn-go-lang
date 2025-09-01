package main

// code segment - all functions - jegula kokhono change hobe na
// data segment - global memory
// stack
// heap

import "fmt"

const p = 100

var a = 10

func add(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println("Sum is:", sum)
}

func call() {
	add := func(x int, y int) {
		z := x + y
		fmt.Println("Add is:", z)
	}

	add(5, 6)
	add(p, a)
}

func main() {
	fmt.Println("Internal Memory in Go lang")

	add(2, 5)

	call()
}

func init() {
	fmt.Println("Welcome")
}
