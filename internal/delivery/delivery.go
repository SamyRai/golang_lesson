package delivery

import (
	"fmt"
)

type Delivery interface {
	PlayerData() (string, string)
	BoardData() int
	MoveData() (int, int)
	Notify(string)
}

type Cli struct {
}

func (c Cli) PlayerData() (string, string) {
	var name string
	var side string

	fmt.Println("Welcome to the game, what's your name and side (x or o)?")
	fmt.Scan(&name, &side)

	return name, side
}

func (c Cli) BoardData() int {
	var boardSize int

	fmt.Println("Please enter a board size")
	fmt.Scan(&boardSize)

	return boardSize
}

func (c Cli) MoveData() (int, int) {
	var xCoordinate int
	var yCoordinate int

	fmt.Println("Enter your next move (x, y coordinates)")
	fmt.Scan(&xCoordinate, &yCoordinate)

	return xCoordinate, yCoordinate
}

func (c Cli) Notify(message string) {
	fmt.Println(message)
}
