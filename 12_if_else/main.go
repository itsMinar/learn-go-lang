package main

import "fmt"

func main() {
	fmt.Println("If Else in Go")

	loginCount := 23

	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount >= 10 && loginCount < 20 {
		result = "Frequent User"
	} else if loginCount >= 20 && loginCount < 30 {
		result = "Very Frequent User"
	} else {
		result = "Super User"
	}

	fmt.Println("Login Count:", loginCount)
	fmt.Println("Result:", result)

	if 9%2 == 0 {
		fmt.Println("9 is even")
	} else {
		fmt.Println("9 is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is 10 or more")
	}
}
