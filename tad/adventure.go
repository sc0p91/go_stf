package main

import (
	"bufio"
	"fmt"
	"github.com/sc0p91/tad/game"
	"os"
	"os/exec"
	"runtime"
)

var clr map[string]func()

func init() {
	clr = make(map[string]func())
	clr["linux"] = func() {
		// Linux Clear
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clr["windows"] = func() {
		// Windows Clear
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func clear() {
	value, ok := clr[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func gameLoop() {

	player := game.NewPlayer()
	enemy := game.NewEnemy(player.Lvl)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() && player.Alive {

		// Clear the screen for every iteration
		clear()

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

		player.Menu()
	}
}

func main() {

	gameLoop()
}
