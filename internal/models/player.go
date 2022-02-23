package models

import (
	"errors"
	"fmt"
	"strings"
)

const WrongSideSymbolError string = "wrong symbol, please pick 'x' or 'o'"
const ShortNameError string = "your name is way too short"

type Player struct {
	name string
	side string
}

func NewPlayer(name string, side string) (Player, error) {
	side = strings.ToLower(side)

	if side != "o" && side != "x" {
		return Player{}, errors.New(WrongSideSymbolError)
	}

	if len(name) <= 2 {
		return Player{}, errors.New(ShortNameError)
	}
	return Player{name, side}, nil
}

func PreparePlayer() (Player, error) {
	for {
		var name string
		var side string

		fmt.Println("Welcome to the game, what's your name and side (x or o)?")
		fmt.Scan(&name, &side)
		
		player, err := NewPlayer(name, side)
		if err == nil {
			return player, nil
		} else {
			fmt.Println(err, "\nTry again!")
		}
	}
}
