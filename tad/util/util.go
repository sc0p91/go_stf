package util

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var clr map[string]func()

func init() {
	clr = make(map[string]func())
	clr["linux"] = func() {
		// Linux Clear
		fmt.Println("\033[2J")
	}
	clr["windows"] = func() {
		// Windows Clear
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func Clear() {
	value, ok := clr[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

