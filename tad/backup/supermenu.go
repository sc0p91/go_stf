package main

import (
  "bufio"
  "fmt"
  "os"
  "time"
  //"strings"
)

func main() {
//  reader := bufio.NewReader(os.Stdin	)
//
//	fmt.Print("Enter Name: ")
//	pname, _ := reader.ReadString('\n')
//	pname = strings.Replace(pname, "\n", "", -1)
//
//	fmt.Print("Enter Class: ")
//	pclass, _ := reader.ReadString('\n')
//
//	fmt.Print("Hi ", pname, " almighty ", pclass)

scanner := bufio.NewScanner(os.Stdin)

for scanner.Scan() {
	fmt.Println("\033[2J")

	text := scanner.Text()

	fmt.Println("- M E N U -")
	fmt.Print("[j] joke\n",
	"[s] start\n",
	"[q] quit\n\n")

	switch text {	
	case "joke","j":
		fmt.Println("This Program.")
	  	break
	case "start","s":
		fmt.Print("starting\n[")
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Print("#")
		}
		fmt.Println("]\ndone")
	  	break
	case "quit","q":
		fmt.Println("quitting")
		os.Exit(0)
	default:
		fmt.Println(text)
		break
	}
}

// reader := bufio.NewReader(os.Stdin)
// char, _, err := reader.ReadRune()
// 
// if err != nil {
//   fmt.Println(err)
// }
// 
// // print out the unicode value i.e. A -> 65, a -> 97
// fmt.Println(char)
// 
// switch char {
// case 'A':
//   fmt.Println("A Key Pressed")
//   break
// case 'a':
//   fmt.Println("a Key Pressed")
//   break
// }

}