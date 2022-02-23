package models

import (
	"errors"
	"fmt"
)

const zeroSizeBoardError = "the board size should be more than zero"
const moveOutsideRangeError = "your move is outside the board range"

type GameBoard struct {
	size  int
	board [][]string
}

func NewGameBoard(size int) (GameBoard, error) {
	if size == 0 {
		return GameBoard{}, errors.New(zeroSizeBoardError)
	}
	return GameBoard{size: size}, nil
}

func PrepareBoard() (GameBoard, error) {
	for {
		var boardSize int
		fmt.Println("Please enter a board size")
		fmt.Scan(&boardSize)
		board, err := NewGameBoard(boardSize)
		if err == nil {
			return board, nil
		} else {
			fmt.Println(err, "\n Try Again!")
		}
	}
}

func (g *GameBoard) MakeMove(xCoordinate int, yCoordinate int, player Player) error {
	if xCoordinate <= g.size || yCoordinate <= g.size {
		return errors.New(moveOutsideRangeError)
	}
	g.board[xCoordinate+1][yCoordinate+1] = player.side
	return nil
}

func (g GameBoard) ShowBoard() {
	//for row := range g.board {
	//	fmt.Printf("%v", row)
	//}
	//	TODO: Make board printer
}
