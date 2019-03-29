package main

import (
	"bufio"
	"fmt"
	"github.com/sc0p91/tad/game"
	"github.com/sc0p91/tad/util"
	"os"
)

func gameLoop() {

	player := game.NewPlayer()
	enemy := game.NewEnemy(player.Lvl)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() && player.Alive {

		// Clear the screen for every iteration
		util.Clear()

		action := scanner.Text()
		switch action {
		case "q", "Q":
			player.Battle(0, enemy)
		case "w", "W":
			player.Battle(1, enemy)
		case "e", "E":
			player.Battle(2, enemy)
		case "r", "R":
			player.Battle(3, enemy)
		case "a", "A":
			fmt.Println("drop ", player.Items[0].Name, "\n ")
			player.UnEquipItem(0)
		case "s", "S":
			fmt.Println("drop ", player.Items[1].Name, "\n ")
			player.UnEquipItem(1)
		case "d", "D":
			player.UsePotion()
		case "f", "F":
			fmt.Println("\n...zzz ZZZ zzz...")
		case "x", "X":
			fmt.Println("Bye")
			os.Exit(0)
		case "":
			fmt.Println("\n ")
		default:
			fmt.Print("\n\"", action, "\" not recognized. Please choose a provided Action\n")
		}

		if !enemy.Alive && player.Alive {
			fmt.Printf("You killed %s, you gained %d Exp.\n", enemy.Name, enemy.Exp)
			player.GainExp(enemy.Exp)

			enemy = game.NewEnemy(player.Lvl)
		}
		// Show E & P Stats + Actions
		fmt.Println("\n‗========= OPPONENT =========‗")
		enemy.ShowStats()

		fmt.Println("‗========== PLAYER ==========‗")
		player.ShowStats()

		util.Menu(*player)
	}
}

func main() {

	gameLoop()
}
