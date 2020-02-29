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
	fmt.Println("Welcome brave Adventurer!\nWhat's your Name?")
	reader := bufio.NewReader(os.Stdin)
	pname, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Set Player Class
	// stats stuff

	// Calculate Player Stats
	p.Name = strings.Replace(pname, "\n", "", -1)

	fmt.Print("[Hit enter to begin your journey]")

	return &p
}
