package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops in Go!")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for i, day := range days {
	// 	fmt.Printf("Index: %v, Day: %v\n", i, day)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			rougueValue++
			break
		}

		fmt.Println("Value is:", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("You are out of the loop now!")
}
