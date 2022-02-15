package main

import (
	"fmt"
	models2 "golang_lesson/internal/models"
)

func main() {

	player, err := models2.PreparePlayer()
	if err != nil {
		panic("something wierd happened")
	}
	board, err := models2.PrepareBoard()
	if err != nil {
		panic("something wierd happened")
	}
	fmt.Printf("%v", player)
	fmt.Printf("%v", board)
}
