package models

import (
	"errors"
	"fmt"
)

const zeroSizeBoardError = "the board size should be more than zero"

type GameBoard struct {
	size  int
	board [][]bool
}

func NewGameBoard(size int) (GameBoard, error) {
	if size == 0 {
		return GameBoard{}, errors.New(zeroSizeBoardError)
	}
	return GameBoard{size: size}, nil
}

func (g GameBoard) MakeMove(xCoordinate int, yCoordinate int) error {
	fmt.Printf("your move was to %d and %d", xCoordinate, yCoordinate)
	return nil
}

func (g GameBoard) ShowBoard() {
	for row := range g.board {
		fmt.Printf("%v", row)
	}
}
