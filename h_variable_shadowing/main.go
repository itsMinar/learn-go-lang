package main

import "fmt"

var a = 10

func main() {
	fmt.Println("Variabel Shadowing in Go lang")

	a := 30

	if a > 18 {
		a := 40
		fmt.Println(a)
	}

	fmt.Println(a)
}
