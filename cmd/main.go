package main

import (
	"fmt"
	"golang_lesson/internal/delivery"
	"golang_lesson/internal/models"
)

func main() {
	cli := new(delivery.Cli)
	var player models.Player
	var board models.GameBoard
	var err error

	for {
		player, err = models.NewPlayer(cli)
		if err != nil {
			cli.Notify(err.Error())
		} else {
			break
		}
	}
	for {
		board, err = models.NewGameBoard(cli)
		if err != nil {
			cli.Notify(err.Error())
		} else {
			break
		}
	}
	fmt.Printf("%v", player)
	fmt.Printf("%v", board)
}
