package main

import (
	"fmt"
	"golang_lesson/models"
)

func main() {

	player, err := models.PreparePlayer()
	if err != nil {
		panic("something wierd happened")
	}

	fmt.Printf("%v", player)
}
