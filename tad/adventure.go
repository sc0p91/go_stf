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
			player.battle(0, enemy)
		case "w", "W":
			player.battle(1, enemy)
		case "e", "E":
			player.battle(2, enemy)
		case "r", "R":
			player.battle(3, enemy)
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

		if !enemy.alive && player.alive {
			fmt.Printf("You killed %s, you gained %d exp.\n", enemy.name, enemy.exp)
			player.gainExp(enemy.exp)

			enemy = newEnemy(player.lvl)
		}
		// Show E & P Stats + Actions
		fmt.Println("\n‗========= OPPONENT =========‗")
		enemy.showStats()

		fmt.Println("‗========== PLAYER ==========‗")
		player.showStats()

		player.menu()
	}
}

func (p entity) menu() {

	for i := 0; i < 4; i++ {
		if p.attacks[i].name == "" {
			p.attacks[i].name = "Not unlocked"
		}
	}

	fmt.Print(
		"‗========== ACTION ==========‗\n",
		" [Q] ", p.attacks[0].name, " (MP: ", p.attacks[0].cost, ") \n",
		" [W] ", p.attacks[1].name, " (MP: ", p.attacks[1].cost, ") \n",
		" [E] ", p.attacks[2].name, " (MP: ", p.attacks[2].cost, ") \n",
		" [R] ", p.attacks[3].name, " (MP: ", p.attacks[3].cost, ") \n",
		" [A] Placeholder Item1\n",
		" [S] Placeholder Item2\n",
		" [D] Restore HP|MP    \n",
		" [F] Do nothing.      \n",
		" [X] quit             \n",
		"≡‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗≡\n",
	)
}

func main() {

	gameLoop()
}
