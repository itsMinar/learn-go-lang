package main

import "fmt"

// struck - custom type
type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Welcome to Struct in Go Lang")

	var user1 User

	user1 = User{
		Name: "Minar",
		Age:  24,
	}

	fmt.Println(user1.Name)
	fmt.Println(user1.Age)

	user2 := User{
		Name: "Habib",
		Age:  28,
	}

	fmt.Println(user2.Name)
	fmt.Println(user2.Age)
}
