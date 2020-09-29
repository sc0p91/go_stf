package main

import (
	"fmt"
	"math/rand"
	"time"
)

func newBoard() [3][3]int {
	var board [3][3]int

	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			board[i][j] = 0
		}
	}
	return board
}

func formater(board [3][3]int) {
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			fmt.Print(board[i][j], " ")
		}
		fmt.Print("\n")
	}
}

func move(board [3][3]int) {
	seed := rand.NewSource(time.Now().UnixNano())
	nran := rand.New(seed)
	moves := 1
	active := true

	for active {
		x := nran.Intn(3)
		y := nran.Intn(3)

		if board[x][y] == 0 {
			if moves%2 == 0 {
				board[x][y] = 1
			} else {
				board[x][y] = 2
			}

			moves += 1

			if moves == 10 {
				active = false
			}
		}
	}

	formater(board)
}

func main() {
	board := newBoard()
	formater(board)

	move(board)
}
