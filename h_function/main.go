package main

import "fmt"

func add(num1 int, num2 int) int {
	sum := num1 + num2

	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2

	return sum, mul
}

func main() {
	fmt.Println("Welcome to Function in Go Lang")

	result := add(2, 5)
	fmt.Println(result)

	num1, num2 := getNumbers(2, 5)
	fmt.Println(num1, num2)

	// get user name as input
	var name string
	fmt.Println("Please enter your name:")
	fmt.Scanln(&name)
	fmt.Println("Hello,", name)
}
