package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// VARIABLES
var pname string
var pclass string

func menu() {
	clear()
	fmt.Print(
		"Name\t", pname, "\n",
		"Class\t", pclass, "\n",
		"Level\t", "Placeholder lvl \n",
		"\n",
		"HP\t", "Placeholder HP \n",
		"MP\t", "Placeholder MP \n",
		"DEX\t", "Placeholder DEX \n",
		"EXP\t", "Placeholder EXP \n",
		"\n",
		".-=============================-.\n",
		"| [Q] Placeholder Atk1          |\n",
		"| [W] Placeholder Atk2          |\n",
		"| [E] Placeholder Atk3          |\n",
		"| [R] Placeholder Atk4          |\n",
		"| [A] Placeholder Item1         |\n",
		"| [S] Placeholder Item2         |\n",
		"| [D] Restore HP|MP             |\n",
		"| [F] Do nothing.               |\n",
		"| [X] quit                      |\n",
		"\\_______________________________/\n")
}

func clear() {
	fmt.Println("\033[2J")
}

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
			break
		case "m", "M", "Mage", "mage":
			pclass = "Mage"
			break
		case "p", "P", "Paladin", "paladin":
			pclass = "Paladin"
			break
		case "r", "R", "Ranger", "ranger":
			pclass = "Ranger"
			break
		default:
			fmt.Print("\"", cclass, "\" not recognized.\n",
				"Please choose a provided class\n")
			break
		}
	}

	fmt.Print("A ", pclass, " it is!\n")
}

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		text := scanner.Text()
		switch text {
		case "q", "Q":
			fmt.Println("This Program.")
			break
		case "w", "W":
			fmt.Println("This Program.")
			break
		case "e", "E":
			fmt.Println("This Program.")
			break
		case "r", "R":
			fmt.Println("This Program.")
			break
		case "a", "A":
			fmt.Println("This Program.")
			break
		case "s", "S":
			fmt.Println("This Program.")
			break
		case "d", "D":
			fmt.Println("This Program.")
			break
		case "f", "F":
			fmt.Println("This Program.")
			break
		case "x", "X":
			fmt.Println("Bye")
			os.Exit(0)
		default:
			fmt.Println(text)
			break
		}
	}
}

func main() {

	initialize()
	menu()
	//scanner()

}
