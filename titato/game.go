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
			board[i][j] = 9
		}
	}
	return board
}

func formater(board [3][3]int) {
	fmt.Print("\n")
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

		if board[x][y] == 9 {
			if moves%2 == 0 {
				board[x][y] = 1
			} else {
				board[x][y] = 2
			}

			moves = moves + checkWin(board)
			if moves >= 10 {
				active = false
			}
		}
	}

	formater(board)
}

func checkWin(board [3][3]int) int {
	active := 1
	if (board[0][0] + board[1][1] + board[2][2]) < 9 {
		// Check Horizontals
		if board[0][0] == board[0][1] && board[0][0] == board[0][2] {
			fmt.Println("\nWINRAR h1")
			active = 9
		} else if board[1][0] == board[1][1] && board[1][0] == board[1][2] {
			fmt.Println("\nWINRAR h2")
			active = 9
		} else if board[2][0] == board[2][1] && board[2][0] == board[2][2] {
			fmt.Println("\nWINRAR h3")
			active = 9
		}
		// Check Verticals
		if board[0][0] == board[1][0] && board[0][0] == board[2][0] {
			fmt.Println("\nWINRAR v1")
			active = 9
		} else if board[0][1] == board[1][1] && board[0][1] == board[2][1] {
			fmt.Println("\nWINRAR v2")
			active = 9
		} else if board[0][2] == board[1][2] && board[0][2] == board[2][2] {
			fmt.Println("\nWINRAR v3")
			active = 9
		}
		// Check Diagonals
		if board[0][0] == board[1][1] && board[0][0] == board[2][2] {
			fmt.Println("\nWINRAR d1")
			active = 9
		} else if board[0][2] == board[1][1] && board[0][2] == board[2][0] {
			fmt.Println("\nWINRAR d2")
			active = 9
		}
	}

	return active
}

func main() {
	board := newBoard()
	formater(board)

	move(board)
}
