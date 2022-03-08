package models

import (
	"errors"
	"golang_lesson/internal/delivery/common"
	"strings"
)

var ErrWrongSideSymbolError error = errors.New("wrong symbol, please pick 'x' or 'o'")
var ErrShortNameError error = errors.New("your name is way too short")

type Player struct {
	name string
	side string
}

func NewPlayer(delivery common.Delivery) (Player, error) {
	name, side := delivery.PlayerData()
	side = strings.ToLower(side)

	if side != "o" && side != "x" {
		return Player{}, ErrWrongSideSymbolError
	}

	if len(name) <= 2 {
		return Player{}, ErrShortNameError
	}
	return Player{name, side}, nil
}

func (p Player) Side() string {
	return p.side
}
