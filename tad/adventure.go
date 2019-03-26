package main

import (
	"bufio"
	"fmt"
	"os"
)

func clear() {
	fmt.Println("\033[2J")
	fmt.Println("")
}

func gameLoop() {

	player := newPlayer()
	enemy := newEnemy(player.lvl)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() && player.alive {
		
		// Clear the screen for every iteration
		clear()

		action := scanner.Text()
		switch action {
		case "q", "Q":
			player.battle("Unarmed Strike", enemy)
		case "w", "W":
			fmt.Println("This Program.")
		case "e", "E":
			fmt.Println("This Program.")
		case "r", "R":
			fmt.Println("This Program.")
		case "a", "A":
			fmt.Println("This Program.")
		case "s", "S":
			fmt.Println("This Program.")
		case "d", "D":
			fmt.Println("This Program.")
		case "f", "F":
			fmt.Println("This Program.")
		case "x", "X":
			fmt.Println("Bye")
			os.Exit(0)
		case "":
		default:
			fmt.Print("\"", action, "\" not recognized. Please choose a provided Action\n")
		}

		if !enemy.alive {
			fmt.Printf("You killed %s, you gained %d exp.\n", enemy.name, enemy.exp)
			player.gainExp(enemy.exp)

			enemy = newEnemy(player.lvl)		
		}
		// Show E & P Stats + Actions
		fmt.Println("\n‗========= OPPONENT =========‗")
		enemy.showStats()


		fmt.Println("‗========== PLAYER ==========‗")
		player.showStats()

		menu(player.attacks)
	}
}

func menu(atk map[string]int) {
	fmt.Print(
		"‗========== ACTION ==========‗\n",
		" [Q] Unarmed Strike \n",
		" [W] ", atk, "\n",
		" [E] Placeholder Atk3 \n",
		" [R] Placeholder Atk4 \n",
		" [A] Placeholder Item1\n",
		" [S] Placeholder Item2\n",
		" [D] Restore HP|MP    \n",
		" [F] Do nothing.      \n",
		" [X] quit             \n",
		"≡‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗≡\n")
}

func main() {

	gameLoop()
}
