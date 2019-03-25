package main

import (
	"fmt"
	"math/rand"
	"time"
)

type enemy struct {
	alive	bool
	name    string
	class   string
	hp      int
	maxHp   int
	mp      int
	maxMp   int
	exp     int
	lvl     int
	sta     int
	str     int
	int     int
	dex     int
	attacks [4]string
	items   [2]string
}

func (e *enemy) selectClass(ename string, cname string) {
	e.name = ename
	e.class = cname
}

func newEnemy(scale int) *enemy {

	// Set Random Seed
	rand.Seed(time.Now().UnixNano())

	// Create random Stats
	var ranA = 0
	var ranB = 0
	var ranC = 0
	var ranD = 0
	for (ranD <= 1) {
		ranA = rand.Intn(8) + 1
		ranB = rand.Intn(8) + 1
		ranC = rand.Intn(8) + 1
		ranD = (20 + scale) - ranA - ranB -ranC
	}

	// Create random HP + MaxHP
	fhp := ((rand.Intn(10) + (8 * scale)) * 10 )

	e := &enemy{
		maxHp:  fhp,
		hp:     fhp,
		mp:     100,
		maxMp:  100,
	}

	e.sta = ranD
	e.str = ranC
	e.dex = ranB
	e.int = ranA
	e.lvl = scale
	e.exp = 25 * scale

	// Set Enemy
	class := rand.Intn(5)

	switch class {
	case 0:
		e.selectClass("Soldier", "Humanoid")
		break
	case 1:
		e.selectClass("Ghoul", "Humanoid")
		break
	case 2:
		e.selectClass("Zombie", "Humanoid")
		break
	case 3:
		e.selectClass("Bat", "Animal")
		break
	case 4:
		e.selectClass("Wolf", "Animal")
		break
	default:
		fmt.Print("\"", class, "\" broken random number lol")
		break
	}

	// Enemy gets born
	e.alive = true

	return e
}

func (e *enemy) screen() {
	clear()

	fmt.Print(
		"‗======== CHALLENGER ========‗\n",
		"  Name\t", e.name, "\n",
		"  Class\t", e.class, "\n",
		"  Level\t", e.lvl, "\n",
		"  EXP\t+", e.exp, " \n",
		"  \n",
		"  Hitpoints\t", e.hp, " / ", e.maxHp, "\n",
		"  Manapoints\t", e.mp, " / ", e.maxMp, "\n",
		"  STA ", e.sta, "\t\tSTR\t", e.str, "\n",
		"  DEX ", e.dex, "\t\tINT\t", e.int, "\n",
		"≡‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗≡\n\n")
}
