package main

import (
	"fmt"
)

type User struct {
	Name     string
	Age      int
	Salary   float64
	FavFoods []string
}

func print(numbers *[3]int) {
	fmt.Println(numbers)
}

func printUserInfo(user *User) {
	fmt.Println("User Name:", user.Name)
	fmt.Println("User Age:", user.Age)
	fmt.Println("User Salary:", user.Salary)
	fmt.Println("User FavFoods:", user.FavFoods)
}

func main() {
	fmt.Println("Welcome to Pointers in Go lang")

	// x := 20

	// p := &x // address of x
	// v := *p // value at address p

	// fmt.Println("Address of x is:", p)
	// fmt.Println("Value of x using pointer is:", v)

	arr := [3]int{1, 2, 3}
	print(&arr)

	user := User{
		Name:     "Minar",
		Age:      25,
		Salary:   5000,
		FavFoods: []string{"Pizza", "Burger"},
	}

	printUserInfo(&user)
}
