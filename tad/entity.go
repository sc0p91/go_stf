package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Entity struct {
	Alive  bool
	Player bool
	Name   string
	class
	item
	hp      int
	mp      int
	maxHp   int
	maxMp   int
	exp     int
	maxExp  int
	lvl     int
	stats   map[string]int
	attacks [4]attack
	items   [2]item
	potions int
}

type class struct {
	name    string
	primary string
}

type item struct {
	name  string
	multi int
	smod  string
}

type attack struct {
	name     string
	damage   int
	cost     int
	lvlreq   int
	classreq string
	slot     int
}

func (e *Entity) gainExp(exp int) {

	// Set Random Seed
	rand.Seed(time.Now().UnixNano())

	e.exp += exp
	if e.exp >= e.maxExp {
		fmt.Println("You leveled up!")
		e.levelUp()
	} else {
		if rand.Intn(5) > 0 {
			e.getItem()
			fmt.Println("You found an Item!")
		} else {
			fmt.Println(" ")
		}
	}
}

func (e *Entity) levelUp() {
	e.lvl++

	e.exp = e.exp - e.maxExp
	e.maxExp = e.maxExp * 2

	for s := range e.stats {
		if s != e.primary {
			e.stats[s]++
		} else {
			e.stats[s] += 2
		}
	}

	e.statRecalc()

	e.hp = e.maxHp
	e.mp = e.maxMp

	for _, a := range attackTemplates {
		if a.classreq == e.class.name && e.lvl >= a.lvlreq {
			e.attacks[a.slot] = a
			e.attacks[a.slot].damage = a.damage + e.stats[e.primary] + e.lvl
			e.attacks[a.slot].cost = a.cost + 2*e.lvl
		}
	}

}

func (e *Entity) statRecalc() {
	e.maxHp = ((e.stats["sta"] + 5) * 10) * e.lvl
	e.maxMp = (e.stats["int"] + 2) * 10 * e.lvl

	e.attacks[0].damage = e.stats["str"] + 25
}

func (e *Entity) getItem() {

	// Set Random Seed
	rand.Seed(time.Now().UnixNano())

	slot := 3
	item := rand.Intn(12)
	luck := rand.Intn(3)

	if e.items[0].name == "none" {
		slot = 0
	} else if e.items[1].name == "none" {
		slot = 1
	}

	if luck == 1 && slot != 3 {
		switch item {
		case 0, 1:
			e.items[slot] = itemTemplates["roh"]
		case 2, 3:
			e.items[slot] = itemTemplates["row"]
		case 4, 5:
			e.items[slot] = itemTemplates["rob"]
		case 6, 7:
			e.items[slot] = itemTemplates["rog"]
		case 8:
			e.items[slot] = itemTemplates["noh"]
		case 9:
			e.items[slot] = itemTemplates["now"]
		case 10:
			e.items[slot] = itemTemplates["nob"]
		case 11:
			e.items[slot] = itemTemplates["nog"]
		}
		// Recalculate Stats with item
		e.equpItem(slot)
		e.statRecalc()
	} else {
		e.potions++
	}
}

func (e *Entity) equpItem(slot int) {
	e.stats[e.items[slot].smod] += e.items[slot].multi
	e.statRecalc()
}

func (e *Entity) unequpItem(slot int) {
	e.stats[e.items[slot].smod] -= e.items[slot].multi
	e.items[slot] = itemTemplates["none"]
	e.statRecalc()
}

func (e *Entity) usePotion() {
	if e.potions < 1 {
		fmt.Println("\nNo Potions left... ")
	} else {
		e.potions--

		php := e.hp + (50 * e.lvl) + (e.stats["sta"] * 10)

		if php > e.maxHp {
			fmt.Println("Filled HP to the Max!")
			e.hp = e.maxHp
		} else {
			fmt.Printf("Healed for %d HP\n", php-e.hp)
			e.hp = php
		}

		pmp := e.mp + (25 * e.lvl) + (e.stats["int"] * 10)

		if pmp > e.maxMp {
			fmt.Println("Filled MP to the Max!")
			e.mp = e.maxMp
		} else {
			fmt.Printf("Restored %d MP\n", pmp-e.mp)
			e.mp = pmp
		}
	}
}

func (e Entity) showStats() {
	fmt.Print(
		" Name\t", e.Name, "\n",
		" Class\t", e.class.name, "\n",
		" Level\t", e.lvl, "\n")
	if e.Player {
		fmt.Print(" EXP\t", e.exp, " / ", e.maxExp, "\n")
	} else {
		fmt.Print(" EXP\t", e.exp, "\n")
	}
	fmt.Print(
		" \n",
		" Hitpoints\t", e.hp, " / ", e.maxHp, "\n",
		" Manapoints\t", e.mp, " / ", e.maxMp, "\n",
		" STA ", e.stats["sta"], "\t\tSTR\t", e.stats["str"], "\n",
		" DEX ", e.stats["dex"], "\t\tINT\t", e.stats["int"], "\n")
	if e.Player {
		fmt.Print(" Itemslot 1:\t", e.items[0].name, "\n")
		fmt.Print(" Itemslot 2:\t", e.items[1].name, "\n")
	}
	fmt.Print("≡‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗≡\n\n")
}

func (e *Entity) battle(attack int, other *Entity) {

	// Set Random Seed
	rand.Seed(time.Now().UnixNano())

	if e.mp >= e.attacks[attack].cost {
		pdmg := e.attacks[attack].damage + rand.Intn(e.lvl)
		dice := rand.Intn(20) + 1

		if other.hp > 0 && pdmg < other.hp {
			if dice == 20 {
				fmt.Println("CRITICAL HIT!", pdmg*2, "DMG!")
				other.hp -= pdmg * 2
			} else if dice == 1 {
				fmt.Println("Your attacked missed...")
			} else {
				fmt.Println("You hit the", other.Name, "for", pdmg, "dmg")
				other.hp -= pdmg
			}
		} else {
			other.hp = 0
			other.Alive = false
		}

		e.mp -= e.attacks[attack].cost

	} else {
		fmt.Println("You don't have enough mana")
	}

	if other.Alive {
		edmg := other.attacks[1].damage + rand.Intn(e.lvl)

		if edmg >= e.hp {
			e.hp = 0
			e.Alive = false
			fmt.Println("You died a horrible death...")
		} else {
			if rand.Intn(10) > 8 {
				if other.mp >= other.attacks[1].cost {
					fmt.Println(other.Name, other.attacks[1].name, "you for ", edmg, "dmg")
					e.hp -= edmg
					other.mp -= other.attacks[1].cost
				} else {
					fmt.Println(other.Name, "misses its attack")
				}
			} else {
				fmt.Println(other.Name, other.attacks[0].name, "you for ", edmg, "dmg")
				e.hp -= edmg
			}
		}
	}
}
