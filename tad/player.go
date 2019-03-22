package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type player struct {
	name    string
	class   string
	hp      uint
	mp      uint
	maxHp   uint
	maxMp   uint
	exp     uint
	maxExp  uint
	lvl     uint
	sta     uint
	str     uint
	int     uint
	dex     uint
	attacks [4]string
	items   [2]string
	hhi     string
	mhi     string
}

func newPlayer() *player {
	p := &player{
		hp:     100,
		mp:     100,
		maxHp:  100,
		maxMp:  100,
		maxExp: 100,
		lvl:    1,
	}
	// Set Player Name
	fmt.Println("Welcome Adventurer!\nWhat's your name?")
	reader := bufio.NewReader(os.Stdin)
	pname, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	p.name = strings.Replace(pname, "\n", "", -1)

	// Set Player Class
	clear()
	fmt.Print("Welcome ", pname, "\n",
		"What Class would you like to be?\n",
		"[W] Warrior - Mighty & Strong Sword Fighter\n",
		"[M] Mage - Powerfull Glasscanon Spellcaster\n",
		"[P] Paladin - Rightful Fighter of the Light\n",
		"[R] Ranger - Distant Fighter with Bow and Arrow\n")

	for p.class == "" {
		class, _ := reader.ReadString('\n')
		class = strings.Replace(class, "\n", "", -1)
		switch class {
		case "w", "W", "Warrior", "warrior":
			fmt.Print("MATCH")
			p.class = "Warrior"
			p.sta = 7
			p.str = 7
			p.dex = 4
			p.int = 2
			break
		case "m", "M", "Mage", "mage":
			p.class = "Mage"
			p.sta = 5
			p.str = 2
			p.dex = 3
			p.int = 10
			break
		case "p", "P", "Paladin", "paladin":
			p.class = "Paladin"
			p.sta = 5
			p.str = 6
			p.dex = 4
			p.int = 5
			break
		case "r", "R", "Ranger", "ranger":
			p.class = "Ranger"
			p.sta = 6
			p.str = 2
			p.dex = 8
			p.int = 4
			break
		default:
			fmt.Print("\"", class, "\" not recognized.\n",
				"Please choose a provided class\n")
			break
		}
	}

	fmt.Print("A ", p.class, " it is!\n")
	return p
}

func (p *player) menu() {
	clear()
	fmt.Print(
		"‗========== PLAYER ===========‗\n",
		"  Name\t", p.name, "\n",
		"  Class\t", p.class, "\n",
		"  Level\t", p.lvl, "\n",
		"  EXP\t", p.exp, " / ", p.maxExp, " \n",
		"  \n",
		"  Hitpoints\t", p.hp, " / ", p.maxHp, "\n",
		"  Manapoints\t", p.mp, " / ", p.maxMp, "\n",
		"  STA ", p.sta, "\t\tSTR\t", p.str, "\n",
		"  DEX ", p.dex, "\t\tINT\t", p.int, "\n",
		"≡‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗≡\n",
		"\n",
		"‗========= ACTIONS ==========‗\n",
		"  [Q] Placeholder Atk1 \n",
		"  [W] Placeholder Atk2 \n",
		"  [E] Placeholder Atk3 \n",
		"  [R] Placeholder Atk4 \n",
		"  [A] Placeholder Item1\n",
		"  [S] Placeholder Item2\n",
		"  [D] Restore HP|MP    \n",
		"  [F] Do nothing.      \n",
		"  [X] quit             \n",
		"≡‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗≡\n")
}
