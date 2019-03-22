package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func initialize() {
	// Set Player Name
	fmt.Println("Welcome Adventurer!\nWhat's your name?")
	reader := bufio.NewReader(os.Stdin)
	pname, _ = reader.ReadString('\n')
	pname = strings.Replace(pname, "\n", "", -1)

	// Set Player Class
	clear()
	fmt.Print("Welcome ", pname, "\n",
		"What Class would you like to be?\n",
		"[W] Warrior - Mighty & Strong Sword Fighter\n",
		"[M] Mage - Powerfull Glasscanon Spellcaster\n",
		"[P] Paladin - Rightful Fighter of the Light\n",
		"[R] Ranger - Distant Fighter with Bow and Arrow\n")

	for pclass == "" {
		cclass, _ := reader.ReadString('\n')
		cclass = strings.Replace(cclass, "\n", "", -1)
		switch cclass {
			case "w", "W", "Warrior", "warrior":
				fmt.Print("MATCH")
				pclass = "Warrior"
				psta = 7
				pstr = 7
				pdex = 4
				pint = 2
				break
			case "m", "M", "Mage", "mage":
				pclass = "Mage"
				psta = 5
				pstr = 2
				pdex = 3
				pint = 10
				break
			case "p", "P", "Paladin", "paladin":
				pclass = "Paladin"
				psta = 5
				pstr = 6
				pdex = 4
				pint = 5
				break
			case "r", "R", "Ranger", "ranger":
				pclass = "Ranger"
				psta = 6
				pstr = 2
				pdex = 8
				pint = 4
				break
			default:
				fmt.Print("\"", cclass, "\" not recognized.\n",
					"Please choose a provided class\n")
				break
		}
	}

	fmt.Print("A ", pclass, " it is!\n")
}
