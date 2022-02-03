package main

import "fmt"

func main() {
	var gameSymbol string
	var currentPlay int

	fmt.Println("Welcome to the Tic-tac-toe game, please pick X or O")

	for gameSymbol != "X" && gameSymbol != "O" {
		fmt.Scanln(&gameSymbol)
		if gameSymbol != "X" && gameSymbol != "O" {
			fmt.Println("Only X and O are allowed")
		}
	}
	fmt.Printf("Awesome, you picked %s, now enter your first move (1..9)\n", gameSymbol)
	fmt.Scanln(&currentPlay)
}
