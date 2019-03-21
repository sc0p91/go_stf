package main

import (
  "bufio"
  "fmt"
  "os"
  "time"
  "strings"
)

func testing() {
	fmt.Print("starting\n[")
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print("#")
	}
	fmt.Println("]\ndone")
}

func menu() {
	fmt.Println("- M E N U -")
	fmt.Print("[j] joke\n",
	"[c] clear\n",
	"[s] start\n",
	"[q] quit\n\n")
}

func clear() {
	fmt.Println("\033[2J")
}

func initialize() {
	pclass := ""
	fmt.Println("Welcome Adventurer!\nWhat's your name?")
	reader := bufio.NewReader(os.Stdin)
	pname, _ := reader.ReadString('\n')
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
			case "w":
				fmt.Print("MATCH")
				pclass = "Warrior"
				break
			case "m","M","Mage","mage":
				pclass = "Mage"
				break
			case "p","P","Paladin","paladin":
				pclass = "Paladin"
				break
			case "r","R","Ranger","ranger":
				pclass = "Ranger"
				break
			default:	
				fmt.Print("No match: ", cclass,
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
		case "joke","j":
			fmt.Println("This Program.")
			  break
		case "start","s":
			testing()
			  break
		case "clear","c":
			clear()
			menu()
			  break
		case "quit","q":
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
	//menu()
	//scanner()

}