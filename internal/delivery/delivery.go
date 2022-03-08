package delivery

import (
	"fmt"
)

type Cli struct {
}

func (c Cli) PlayerData() (name string, side string) {
	fmt.Println("Welcome to the game, what's your name and side (x or o)?")
	fmt.Scan(&name, &side)

	return name, side
}

func (c Cli) BoardData() (boardSize int) {
	fmt.Println("Please enter a board size")
	fmt.Scan(&boardSize)

	return boardSize
}

func (c Cli) MoveData() (xCoordinate int, yCoordinate int) {
	fmt.Println("Enter your next move (x, y coordinates)")
	fmt.Scan(&xCoordinate, &yCoordinate)

	return xCoordinate, yCoordinate
}

func (c Cli) ShowBoard(board [][]string) {
	var stringBoard string

	fmt.Println("Your board looks like this:")
	stringIndex := 0
	for _, row := range board {
		stringBoard += "|"

		for _, col := range row {
			place := col
			if place == "" {
				place = " "
			}
			stringBoard = stringBoard + place + "|"
			stringIndex++
		}
		stringBoard += "\n"
	}
	fmt.Println(stringBoard)
}

func (c Cli) Notify(message string) {
	fmt.Println(message)
}
