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

	player := NewPlayer()
	enemy := NewEnemy(player.Lvl)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() && player.Alive {

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
			fmt.Println("drop ", player.Items[0].Name, "\n ")
			player.unequpItem(0)
		case "s", "S":
			fmt.Println("drop ", player.Items[1].Name, "\n ")
			player.unequpItem(1)
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

		if !enemy.Alive && player.Alive {
			fmt.Printf("You killed %s, you gained %d Exp.\n", enemy.Name, enemy.Exp)
			player.gainExp(enemy.Exp)

			enemy = NewEnemy(player.Lvl)
		}
		// Show E & P Stats + Actions
		fmt.Println("\n‗========= OPPONENT =========‗")
		enemy.showStats()

		fmt.Println("‗========== PLAYER ==========‗")
		player.showStats()

		player.menu()
	}
}

func (p Entity) menu() {

	for i := 0; i < 4; i++ {
		if p.Attacks[i].Name == "" {
			p.Attacks[i].Name = "Not unlocked"
		}
	}

	fmt.Print(
		"‗========== ACTION ==========‗\n",
		" [Q] ", p.Attacks[0].Name, " (MP: ", p.Attacks[0].Cost, ") \n",
		" [W] ", p.Attacks[1].Name, " (MP: ", p.Attacks[1].Cost, ") \n",
		" [E] ", p.Attacks[2].Name, " (MP: ", p.Attacks[2].Cost, ") \n",
		" [R] ", p.Attacks[3].Name, " (MP: ", p.Attacks[3].Cost, ") \n",
		" [A] Drop Item 1 \n",
		" [S] Drop Item 2 \n",
		" [D] Restore HP&MP (#: ", p.Potions, ")\n",
		" [F] Do nothing.      \n",
		" [X] quit             \n",
		"≡‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗≡\n",
	)
}

func main() {

	gameLoop()
}
