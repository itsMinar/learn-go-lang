package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Go")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	languages["GO"] = "Go"

	fmt.Println("Languages:", languages)
	fmt.Println("JS short for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("After deleting RB:", languages)

	// loops are interesting
	for key, value := range languages {
		fmt.Printf("%v is short for %v\n", key, value)
	}
}
