package main

import (
	"fmt"
	"h_scope/mathlib"
)

var a = 20
var b = 30

func main() {
	fmt.Println("Welcome to Scope in Go Lang")

	p := 30
	q := 40

	mathlib.Add(p, q)

	mathlib.Add(a, b)

	mathlib.Add(a, p)

}
