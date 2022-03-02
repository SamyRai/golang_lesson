package game

import (
	"errors"
	"golang_lesson/internal/delivery/common"
	"golang_lesson/internal/models"
)

var ErrMoveOutsideRange = errors.New("your move is outside the board range")
var ErrPlaceIsAlreadyTaken = errors.New("your place was already taken, try another one")

func MakeMove(gameBoard models.GameBoard, delivery common.Delivery, player models.Player) error {
	xCoordinate, yCoordinate := delivery.MoveData()
	if xCoordinate > gameBoard.Size() || yCoordinate > gameBoard.Size() {
		return ErrMoveOutsideRange
	}
	_, err := gameBoard.RecordMove(xCoordinate, yCoordinate, player)

	if gameBoard.IsOccupied(xCoordinate, yCoordinate) {
		return ErrPlaceIsAlreadyTaken
	}

	if err != nil {
		return err
	}

	return nil
}
