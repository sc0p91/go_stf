package util

import (
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
		fmt.Println("\033[2J")
	}
	clr["windows"] = func() {
		// Windows Clear
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func Clear() {
	value, ok := clr[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func Menu(p game.Entity) {

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
