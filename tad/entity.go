package main

import "fmt"

type entity struct {
	alive   bool
	player  bool
	name    string
	class
	hp      int
	mp      int
	maxHp   int
	maxMp   int
	exp     int
	maxExp  int
	lvl     int
	sta     int
	str     int
	int     int
	dex     int
	attacks map[string]int
	items   [2]string
}

type class struct {
	name string
	primary string
}

func (e *entity) selectClass(name string, sta int, str int, dex int, int int, primary string) {
	e.class.name = name
	e.primary = primary
	e.sta = sta
	e.str = str
	e.dex = dex
	e.int = int
}

func (e *entity) gainExp(exp int) {
	e.exp += exp
	if e.exp >= e.maxExp {
		e.levelUp()
	}
}

func (e *entity) levelUp() {
	e.lvl++
	e.exp = e.exp-e.maxExp
	e.maxExp = e.maxExp * 2
	e.sta += 2
	e.hp += e.sta * 10
	e.maxHp += e.sta * 10
	e.mp += e.int * 10
	e.maxMp += e.int * 10

	e.attacks["Unarmed Strike"] = e.str + 25

	switch e.primary {
		case "str":
			e.int++
			e.dex++
			e.str += 2
			e.attacks["Obliterate"] = e.str + 25
		case "int":
			e.int+= 2
			e.dex++
			e.str++
			e.attacks["Lightning Bolt"] = e.int + 25
		case "dex":
			e.int++
			e.dex+= 2
			e.str++
			e.attacks["Cobra Shot"] = e.dex + 25
	}

	fmt.Println("You gained a level")
}

func (e entity) showStats() {
	fmt.Print(
		" Name\t", e.name, "\n",
		" Class\t", e.class.name, "\n",
		" Level\t", e.lvl, "\n")
	if e.player {
		fmt.Print(" EXP\t", e.exp," / ", e.maxExp, "\n")
	} else {
		fmt.Print(" EXP\t", e.exp, "\n")
	}
	fmt.Print(
		" \n",
		" Hitpoints\t", e.hp, " / ", e.maxHp, "\n",
		" Manapoints\t", e.mp, " / ", e.maxMp, "\n",
		" STA ", e.sta, "\t\tSTR\t", e.str, "\n",
		" DEX ", e.dex, "\t\tINT\t", e.int, "\n",
		"≡‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗≡\n\n")
}

func (e *entity) battle(attack string, other *entity) {	
	fmt.Println("You hit the", other.name, "for", e.attacks[attack], "dmg")
	if ( other.hp > 0 && e.attacks[attack] < other.hp ) {
		other.hp -= e.attacks[attack]
	} else {
		other.hp = 0
		other.alive = false
	}
}