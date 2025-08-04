package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Coffee Shop App!")
	fmt.Println("Please rate our Coffee from 1 to 5:")

	reader := bufio.NewReader(os.Stdin)

	// comma or || error syntax
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for your rating, ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Error converting input to number:", err)
	} else {
		fmt.Println("Added 1 to your rating:", numRating+1)
	}

}
