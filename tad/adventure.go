package main

import (
	"bufio"
	"fmt"
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
			fmt.Println("drop ", player.items[0].name, "\n ")
		case "s", "S":
			fmt.Println("drop ", player.items[1].name, "\n ")
		case "d", "D":
			player.usePotion()
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
		" [A] Drop Item 1 \n",
		" [S] Drop Item 2 \n",
		" [D] Restore HP&MP (#: ", p.potions, ")\n",
		" [F] Do nothing.      \n",
		" [X] quit             \n",
		"≡‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗≡\n",
	)
}

func main() {

	gameLoop()
}
