package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func newPlayer() *entity {
	p := &entity{
		hp:     100,
		mp:     100,
		maxHp:  100,
		maxMp:  100,
		maxExp: 100,
		lvl:    1,
		player: true,
		alive:  true,
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
			p.selectClass("Warrior", 7, 7, 4, 2)
			break
		case "m", "M", "Mage", "mage":
			p.selectClass("Mage", 5, 2, 3, 10)
			break
		case "p", "P", "Paladin", "paladin":
			p.selectClass("Palading", 5, 6, 4, 5)
			break
		case "r", "R", "Ranger", "ranger":
			p.selectClass("Ranger", 6, 2, 8, 4)
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

func (p *entity) menu() {
	fmt.Print(
		"‗========== ACTION ==========‗\n",
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
