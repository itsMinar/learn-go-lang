package main

import "fmt"

// pointer, length, capacity
// capacity = 2^n where n is the number of bits required to represent the capacity

func TryAllTypesOfSlices() {
	fmt.Println("Welcome to Slices in Go Lang")

	// array
	arr := [6]string{"This", "is", "a", "Go", "interview", "question"}
	fmt.Println("arr:", arr)

	// slice from array
	s := arr[1:4] // start from 1 and end to (4-1)
	fmt.Println("s:", s)

	//  slice from slice
	s1 := s[1:3]
	fmt.Println("s1:", s1)

	// slice literal
	s2 := []int{1, 2, 3}
	fmt.Println("s2:", s2, "len:", len(s2), "cap:", cap(s2))

	// make function with length
	s3 := make([]int, 3)
	// set value
	s3[0] = 100
	fmt.Println("s3:", s3, "len:", len(s3), "cap:", cap(s3))

	// make function with length and capacity
	s4 := make([]int, 3, 5)
	s4[2] = 200
	fmt.Println("s4:", s4, "len:", len(s4), "cap:", cap(s4))

	var s5 []int // empty slice or nil slice
	s5 = append(s5, 1, 2, 3)
	s5 = append(s5, 2)
	fmt.Println("s5:", s5, "len:", len(s5), "cap:", cap(s5))
}
