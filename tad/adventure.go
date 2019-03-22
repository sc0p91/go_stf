package main

import (
	"bufio"
	"fmt"
	"os"
)

func clear() {
	fmt.Println("\033[2J")
	fmt.Println("")
}

func scanner() {

	player := newPlayer()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		
		player.menu()
		
		action := scanner.Text()
		switch action {
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
			break
		case "":
			break
		default:
			fmt.Print("\"", action, "\" not recognized. Please choose a provided Action\n")
			break
		}
	}
}

func main() {
	
	scanner()
}
