package models

import (
	"errors"
	"golang_lesson/internal/delivery/common"
)

var ErrZeroSizeBoard = errors.New("the board size should be more than zero")

type GameBoard struct {
	size  int
	board [][]string
}

func NewGameBoard(delivery common.Delivery) (GameBoard, error) {
	size := delivery.BoardData()

	if size == 0 {
		return GameBoard{}, ErrZeroSizeBoard
	}
	var board = make([][]string, size, 100)
	for i := 0; i < size; i++ {
		board[i] = make([]string, size, 100)
	}
	return GameBoard{size: size, board: board}, nil
}

func (g GameBoard) Size() int {
	return g.size
}

func (g *GameBoard) RecordMove(xCoordinate int, yCoordinate int, player Player) (GameBoard, error) {
	g.board[xCoordinate-1][yCoordinate-1] = player.Side()

	return *g, nil
}

func (g GameBoard) IsOccupied(xCoordinate int, yCoordinate int) bool {
	return g.board[xCoordinate][yCoordinate] == ""
}

func (g GameBoard) ShowBoard(delivery common.Delivery) {
	delivery.ShowBoard(g.board)
}
