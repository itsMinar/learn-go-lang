package main

import "fmt"

func add(num1 int, num2 int) int {
	sum := num1 + num2

	return sum
}

// assign function in variable
var agun = func() {
	fmt.Println("Valo lage na")
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

/***
* different types of function
1. standard or named function
2. anonymous function
3. function expression or assign function in variable
4. higher order function or first class function
5. callback function
6. variadic function
7. init function - you can't call this, pc call this auto
8. Closure - close over
9. Defer function
10. Receiver function
11. IIFE - Immediately Invoked Function Expression
*/
