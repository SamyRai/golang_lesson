package main

import (
	"fmt"
	"golang_lesson/internal/delivery"
	"golang_lesson/internal/models"
)

func main() {
	input := new(delivery.Cli)
	player, err := models.PreparePlayer()
	if err != nil {
		panic("something wierd happened")
	}
	board, err := models.PrepareBoard()
	if err != nil {
		panic("something wierd happened")
	}
	fmt.Printf("%v", player)
	fmt.Printf("%v", board)
}
