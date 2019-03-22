package main

import (
	"bufio"
	"fmt"
	"os"
)

// VARIABLES
var pname string
var pclass string

func clear() {
	fmt.Println("\033[2J")
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
