package main

import (
	"fmt"

	"tic-tac-toe/tictactoe"
)

func main() {
	game := tictactoe.NewTicTacToe()
	if err := game.Start(); err != nil {
		fmt.Printf("Unexpected error: %s\n", err.Error())
	}
}
