package models

import (
	"errors"
	"golang_lesson/internal/delivery"
)

var ErrZeroSizeBoardError = errors.New("the board size should be more than zero")

type GameBoard struct {
	size  int
	board [][]string
}

func NewGameBoard(delivery delivery.Delivery) (GameBoard, error) {
	size := delivery.BoardData()

	if size == 0 {
		return GameBoard{}, ErrZeroSizeBoardError
	}
	return GameBoard{size: size}, nil
}

func (g GameBoard) Size() int {
	return g.size
}

func (g GameBoard) RecordMove(xCoordinate int, yCoordinate int, symbol string) (GameBoard, error) {
	g.board[xCoordinate+1][yCoordinate+1] = symbol

	return g, nil
}

func (g GameBoard) ShowBoard() {
	//for row := range g.board {
	//	fmt.Printf("%v", row)
	//}
	//	TODO: Make board printer
}
