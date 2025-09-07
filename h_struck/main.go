package main

import "fmt"

// struck - custom type
type User struct {
	Name string
	Age  int
}

// receiver function - only possible for struct
func (usr User) printUserDetails() {
	fmt.Println("User Name is:", usr.Name)
	fmt.Println("User Age is:", usr.Age)
}
func (usr User) call(num int) {
	fmt.Println("User Name is:", usr.Name)
	fmt.Println("User Call is:", num)
}

func main() {
	fmt.Println("Welcome to Struct in Go Lang")

	var user1 User

	user1 = User{
		Name: "Minar",
		Age:  24,
	}

	user1.printUserDetails()
	user1.call(1234567890)

	user2 := User{
		Name: "Habib",
		Age:  28,
	}

	user2.printUserDetails()
	user2.call(1234567890)
}
