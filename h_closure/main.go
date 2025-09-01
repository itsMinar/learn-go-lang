package main

import "fmt"

const a = 10

var p = 100

func outer() func() {
	money := 100
	age := 30

	fmt.Println("Age is:", age)

	show := func() {
		money = money + a + p
	}

	return show
}

func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}

func main() {
	fmt.Println("Closure in Go lang")
}

func init() {
	fmt.Println("Welcome")
}
