package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Banana", "Cherry"}
	fmt.Println("Fruit List:", fruitList)

	fruitList = append(fruitList, "Date", "Elderberry")
	fmt.Println("Updated Fruit List:", fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println("Sliced Fruit List:", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 100
	highScores[1] = 200
	highScores[3] = 300
	highScores[2] = 400
	// highScores[4] = 500
	highScores = append(highScores, 500)
	fmt.Println("High Scores:", highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("Sorted High Scores:", highScores)

	// how to remove an element from a slice
	var courses = []string{"ReactJS", "AngularJS", "VueJS", "NodeJS"}
	fmt.Println("Courses:", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Updated Courses:", courses)

}
