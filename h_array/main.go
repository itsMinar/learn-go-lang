package main

import "fmt"

type Player struct {
	name         string
	age          int
	jerseyNumber int
}

func main() {
	fmt.Println("Welcome to Arrays in Go Lang")

	var arr [2]int
	arr[0] = 100
	arr[1] = 200
	fmt.Println("arr", arr)

	var player_arr [2]Player
	player1 := Player{
		name:         "Minar",
		age:          25,
		jerseyNumber: 10,
	}
	player_arr[0] = player1
	player_arr[1] = Player{
		name:         "Rahim",
		age:          26,
		jerseyNumber: 11,
	}

	fmt.Println("Player list is:", player_arr)

	second_arr := [2]int{2, 50}
	fmt.Println("Second array is:", second_arr)

}
