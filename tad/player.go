package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func newPlayer() *entity {
	p := entity{}

	// Set Player Name
	fmt.Println("Welcome Adventurer!\nWhat's your name?")
	reader := bufio.NewReader(os.Stdin)
	pname, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Set Player Class
	fmt.Print("\nGreetings ", pname, "\n",
		"What Class would you like to be?\n",
		"[W] Warrior - Mighty & Strong Sword Fighter\n",
		"[M] Mage - Powerfull Glasscanon Spellcaster\n",
		"[P] Paladin - Rightful Fighter of the Light\n",
		"[R] Ranger - Distant Fighter with Bow and Arrow\n")

	for p.class.name == "" {
		class, _ := reader.ReadString('\n')
		class = strings.Replace(class, "\n", "", -1)
		switch class {
		case "w", "W", "Warrior", "warrior":
			p = entityTemplates["warrior"]
		case "m", "M", "Mage", "mage":
			p = entityTemplates["mage"]
		case "p", "P", "Paladin", "paladin":
			p = entityTemplates["paladin"]
		case "r", "R", "Ranger", "ranger":
			p = entityTemplates["ranger"]
		default:
			fmt.Print("\"", class, "\" not recognized.\n",
				"Please choose a provided class\n")
			break
		}
	}
	p.name = strings.Replace(pname, "\n", "", -1)
	p.hp = 100 + p.stats["sta"]*10
	p.mp = 100 + p.stats["int"]*10
	p.maxHp = p.hp
	p.maxMp = p.mp
	p.maxExp = 100
	p.lvl = 1
	p.player = true
	p.alive = true
	p.items[0] = itemTemplates["none"]
	p.items[1] = itemTemplates["none"]
	p.potions = 3

	fmt.Print("A ", p.class.name, " it is!\n",
		"[Hit enter to begin your journey]")

	p.attacks[0] = attackTemplates["Unarmed Strike"]

	return &p
}
