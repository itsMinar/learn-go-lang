package main

import (
	"fmt"
	"os"
)

// only signature of the function
type People interface {
	PrintDetails()
	ReceiveMoney(amount float64) float64
}

type BankUser interface {
	WithdrawMoney(amount float64) float64
}

type User struct {
	Name  string
	Age   int
	Money float64
}

// receiver function or method or behaviour
func (obj User) PrintDetails() {
	fmt.Println("Name:", obj.Name)
	fmt.Println("Age:", obj.Age)
	fmt.Println("Money:", obj.Money)
}

func (obj User) ReceiveMoney(amount float64) float64 {
	obj.Money += amount
	return obj.Money
}

func (obj User) WithdrawMoney(amount float64) float64 {
	obj.Money -= amount
	return obj.Money
}

func main() {
	var usr1 People
	usr1 = User{
		Name:  "Minar",
		Age:   25,
		Money: 1000,
	}

	fmt.Println("usr1", usr1)
	usr1.PrintDetails()
	receivedAmount := usr1.ReceiveMoney(5000)
	fmt.Println("receivedAmount", receivedAmount)

	var usr2 BankUser
	usr2 = User{
		Name:  "Agun",
		Age:   22,
		Money: 2000,
	}

	fmt.Println("usr2", usr2)
	withdrawnAmount2 := usr2.WithdrawMoney(50)
	fmt.Println("withdrawnAmount2", withdrawnAmount2)

	obj, ok := usr2.(User)
	if !ok {
		fmt.Println("usr2 is not a User")
		os.Exit(1)
	}

	obj.PrintDetails()
}
