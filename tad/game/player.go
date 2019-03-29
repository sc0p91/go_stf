package game

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func NewPlayer() *Entity {
	p := Entity{}

	// Set Player Name
	fmt.Println("Welcome Adventurer!\nWhat's your Name?")
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

	for p.Class.Name == "" {
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
				"Please choose a provided Class\n")
			break
		}
	}
	p.Name = strings.Replace(pname, "\n", "", -1)
	p.Hp = 100 + p.Stats["sta"]*10
	p.Mp = 100 + p.Stats["int"]*10
	p.MaxHp = p.Hp
	p.MaxMp = p.Mp
	p.MaxExp = 100
	p.Lvl = 1
	p.Player = true
	p.Alive = true
	p.Items[0] = itemTemplates["none"]
	p.Items[1] = itemTemplates["none"]
	p.Potions = 3

	fmt.Print("A ", p.Class.Name, " it is!\n",
		"[Hit enter to begin your journey]")

	p.Attacks[0] = attackTemplates["Unarmed Strike"]

	return &p
}
