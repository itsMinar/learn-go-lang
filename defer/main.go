package main

import "fmt"

func a() {
	i := 0

	fmt.Println("first", i)

	defer fmt.Println("second", i)

	i = i + 1

	fmt.Println("third", i)
	defer fmt.Println("fourth", i)
}

// named return value
func sum(a int, b int) (result int) {
	result = a + b
	return
}

func calculate() (result int) {
	fmt.Println("first", result)

	show := func() {
		result = result + 10
		fmt.Println("defer", result)
	}

	defer show()

	result = 5
	fmt.Println("second", result)
	return
}

func calculate2() int {
	result := 0
	fmt.Println("first", result)

	show := func() {
		result = result + 10
		fmt.Println("defer", result)
	}

	defer show()

	result = 5
	fmt.Println("second", result)
	return result
}

func main() {
	a := calculate()
	fmt.Println("a=> ", a)

	b := calculate2()
	fmt.Println("b=> ", b)
}
