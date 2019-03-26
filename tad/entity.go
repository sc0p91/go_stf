package main

import "fmt"

type entity struct {
	alive  bool
	player bool
	name   string
	class
	hp      int
	mp      int
	maxHp   int
	maxMp   int
	exp     int
	maxExp  int
	lvl     int
	stats   map[string]int
	attacks [4]attack
	items   [2]string
}

type class struct {
	name    string
	primary string
}

type attack struct {
	name     string
	damage   int
	cost     int
	lvlreq   int
	classreq class
	slot     int
}

func (e *entity) gainExp(exp int) {
	e.exp += exp
	if e.exp >= e.maxExp {
		e.levelUp()
	}
}

func (e *entity) levelUp() {
	e.lvl++
	e.exp = e.exp - e.maxExp
	e.maxExp = e.maxExp * 2
	e.stats["sta"] += 2
	e.maxHp += e.stats["sta"] * 10
	e.hp = e.maxHp
	e.maxMp += e.stats["int"] * 10
	e.mp = e.maxMp

	e.attacks[0].damage = e.stats["str"] + 25

	for s := range e.stats {
		if s != e.primary {
			e.stats[s]++
		} else {
			e.stats[s] += 2
		}
	}

	for _, a := range attackTemplates {
		if a.classreq.name == e.class.name && e.lvl >= a.lvlreq {
			e.attacks[a.slot] = a
			e.attacks[a.slot].damage = a.damage + e.stats[e.primary]
			e.attacks[a.slot].cost = a.cost + 2*e.lvl
		}
	}

	fmt.Println("You gained a level")
}

func (e entity) showStats() {
	fmt.Print(
		" Name\t", e.name, "\n",
		" Class\t", e.class.name, "\n",
		" Level\t", e.lvl, "\n")
	if e.player {
		fmt.Print(" EXP\t", e.exp, " / ", e.maxExp, "\n")
	} else {
		fmt.Print(" EXP\t", e.exp, "\n")
	}
	fmt.Print(
		" \n",
		" Hitpoints\t", e.hp, " / ", e.maxHp, "\n",
		" Manapoints\t", e.mp, " / ", e.maxMp, "\n",
		" STA ", e.stats["sta"], "\t\tSTR\t", e.stats["str"], "\n",
		" DEX ", e.stats["dex"], "\t\tINT\t", e.stats["int"], "\n",
		"≡‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗≡\n\n")
}

func (e *entity) battle(attack int, other *entity) {
	if e.mp >= e.attacks[attack].cost {
		if other.hp > 0 && e.attacks[attack].damage < other.hp {
			fmt.Println("You hit the", other.name, "for", e.attacks[attack].damage, "dmg")
			other.hp -= e.attacks[attack].damage
		} else {
			other.hp = 0
			other.alive = false
		}
		e.mp -= e.attacks[attack].cost
	} else {
		fmt.Println("You don't have enough mana")
	}
	if other.alive {
		fmt.Println(other.name, other.attacks[0].name, "you for ", other.attacks[0].damage, "dmg")
		if other.attacks[0].damage >= e.hp {
			e.hp = 0
			e.alive = false
			fmt.Println("you dead lol")
		} else {
			e.hp -= other.attacks[0].damage
		}
	}
}
