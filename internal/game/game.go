package game

import (
	"errors"
	"golang_lesson/internal/models"
)

var ErrMoveOutsideRangeError = errors.New("your move is outside the board range")

func MakeMove(gameBoard models.GameBoard, xCoordinate int, yCoordinate int, player models.Player) error {
	if xCoordinate <= gameBoard.Size() || yCoordinate <= gameBoard.Size() {
		return ErrMoveOutsideRangeError
	}
	gameBoard.RecordMove(xCoordinate, yCoordinate, player.Side())
	return nil
}
