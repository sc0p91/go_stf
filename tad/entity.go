package main

import "fmt"

type entity struct {
	alive   bool
	player  bool
	name    string
	class   string
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
	attacks [4]string
	items   [2]string
}

func (e *entity) selectClass(name string, sta int, str int, dex int, int int) {
	e.class = name
	e.sta = sta
	e.str = str
	e.dex = dex
	e.int = int
}

func (e entity) showStats() {
	fmt.Print(
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
