package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case in Go Lang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of Dice is:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You rolled a one!")
	case 2:
		fmt.Println("You rolled a two!")
	case 3:
		fmt.Println("You rolled a three!")
		fallthrough
	case 4:
		fmt.Println("You rolled a four!")
		fallthrough
	case 5:
		fmt.Println("You rolled a five!")
	case 6:
		fmt.Println("You rolled a six!")
	default:
		fmt.Println("What was that?")
	}
}
