package main

import "fmt"

func main() {
	fmt.Println("Welcome to Mehtods in Go!")

	minar := User{"Minar", "minar@example.com", true, 26}
	fmt.Println(minar)
	fmt.Printf("Name: %s, Email: %s, Status: %t, Age: %d\n", minar.Name, minar.Email, minar.Status, minar.Age)
	fmt.Printf("Type of minar: %T\n", minar)

	minar.GetName()
	minar.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetName() {
	fmt.Println("Is User active?", u.Status)

}

func (u User) NewMail() {
	u.Email = "minar@example.com"
	fmt.Println("New Mail Sent to", u.Email)
}
