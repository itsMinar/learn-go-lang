package main

import "fmt"

const LoginToken string = "asdasdasdasdasdasdasdasdasdasd"

func main() {
	var username string = "Minar"

	fmt.Println("Hello, " + username + "!")
	fmt.Printf("Type of username is: %T\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Value of isLoggedIn is: %v\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Value of smallVal is: %v\n", smallVal)

	// no var style
	numberofUsers := 3000.0
	fmt.Println(numberofUsers)
	fmt.Printf("Type of numberofUsers is: %T\n", numberofUsers)

	fmt.Println("Login token is:", LoginToken)
}
