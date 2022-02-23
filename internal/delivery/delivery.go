package delivery

import (
	"fmt"
	"golang_lesson/internal/models"
)

type delivery interface {
	playerData() (models.Player, error)
	boardData() (models.GameBoard, error)
	moveData() (models.GameBoard, error)
}

type Cli struct {
}

func (c Cli) playerData() (models.Player, error) {
	var name string
	var side string

	fmt.Println("Welcome to the game, what's your name and side (x or o)?")
	fmt.Scan(&name, &side)

	player, err := models.NewPlayer(name, side)
	if err == nil {
		return player, nil
	} else {
		fmt.Println(err, "\nTry again!")
	}
}

func (c Cli) boardData() (models.GameBoard, error) {
	for {
		var boardSize int
		fmt.Println("Please enter a board size")
		fmt.Scan(&boardSize)
		board, err := models.NewGameBoard(boardSize)
		if err == nil {
			return board, nil
		} else {
			fmt.Println(err, "\n Try Again!")
		}
	}
}
