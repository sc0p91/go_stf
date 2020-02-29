package main

import (
	"bufio"
	"fmt"
	"github.com/sc0p91/tad2/game"
	"github.com/sc0p91/tad2/util"
	"os"
)

func gameLoop() {

	player := game.NewPlayer()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() && player.Alive {

		// Clear the screen for every iteration
		util.Clear()

		action := scanner.Text()
		switch action {
		case "q", "Q":
			player.Battle(0, enemy)
		case "a", "A":
			fmt.Println("drop ", player.Items[0].Name, "\n ")
			player.UnEquipItem(0)
		case "d", "D":
			player.UsePotion()
		case "x", "X":
			fmt.Println("Bye")
			os.Exit(0)
		case "":
			fmt.Println("\n ")
		default:
			fmt.Print("\n\"", action, "\" not recognized. Please choose a provided Action\n")
		}
	}
}

func main() {

	gameLoop()
}
