package main

import "fmt"

func main() {
	fmt.Println("===>Structs in Go<===")
	// no inheritance, no super or parent

	minar := User{"Minar", "minar@example.com", true, 26}
	fmt.Println(minar)
	fmt.Printf("Name: %s, Email: %s, Status: %t, Age: %d\n", minar.Name, minar.Email, minar.Status, minar.Age)
	fmt.Printf("Type of minar: %T\n", minar)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
