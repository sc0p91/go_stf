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
	fmt.Print("\nGreetings ", pname, "\n",
		"What Class would you like to be?\n",
		"[W] Warrior - Mighty & Strong Sword Fighter\n",
		"[M] Mage - Powerfull Glasscanon Spellcaster\n",
		"[P] Paladin - Rightful Fighter rof the Light\n",
		"[R] Ranger - Distant Fighter with Bow and Arrow\n")

	for p.class.name == "" {
		class, _ := reader.ReadString('\n')
		class = strings.Replace(class, "\n", "", -1)
		switch class {
		case "w", "W", "Warrior", "warrior":
			p.selectClass("Warrior", 7, 7, 4, 2, "str")
			break
		case "m", "M", "Mage", "mage":
			p.selectClass("Mage", 5, 2, 3, 10, "int")
			break
		case "p", "P", "Paladin", "paladin":
			p.selectClass("Palading", 5, 6, 4, 5, "int")
			break
		case "r", "R", "Ranger", "ranger":
			p.selectClass("Ranger", 6, 2, 8, 4, "dex")
			break
		default:
			fmt.Print("\"", class, "\" not recognized.\n",
				"Please choose a provided class\n")
			break
		}
	}
	fmt.Print("A ", p.class.name, " it is!\n",
		"[Hit enter to begin your journey]")

	p.attacks[0] = attack{"Unarmed Strike", p.str + 25, 10}

	return p
}
