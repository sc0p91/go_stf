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
	Class
	Item
	Hp      int
	Mp      int
	MaxHp   int
	MaxMp   int
	Exp     int
	MaxExp  int
	Lvl     int
	Stats   map[string]int
	Attacks [4]Attack
	Items   [2]Item
	Potions int
}

type Class struct {
	Name    string
	Primary string
}

type Item struct {
	Name  string
	Multi int
	Smod  string
}

type Attack struct {
	Name     string
	Damage   int
	Cost     int
	LvlReq   int
	ClassReq string
	Slot     int
}

func (e *Entity) gainExp(exp int) {

	// Set Random Seed
	rand.Seed(time.Now().UnixNano())

	e.Exp += exp
	if e.Exp >= e.MaxExp {
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
	e.Lvl++

	e.Exp = e.Exp - e.MaxExp
	e.MaxExp = e.MaxExp * 2

	for s := range e.Stats {
		if s != e.Primary {
			e.Stats[s]++
		} else {
			e.Stats[s] += 2
		}
	}

	e.statRecalc()

	e.Hp = e.MaxHp
	e.Mp = e.MaxMp

	for _, a := range attackTemplates {
		if a.ClassReq == e.Class.Name && e.Lvl >= a.LvlReq {
			e.Attacks[a.Slot] = a
			e.Attacks[a.Slot].Damage = a.Damage + e.Stats[e.Primary] + e.Lvl
			e.Attacks[a.Slot].Cost = a.Cost + 2*e.Lvl
		}
	}

}

func (e *Entity) statRecalc() {
	e.MaxHp = ((e.Stats["sta"] + 5) * 10) * e.Lvl
	e.MaxMp = (e.Stats["int"] + 2) * 10 * e.Lvl

	e.Attacks[0].Damage = e.Stats["str"] + 25
}

func (e *Entity) getItem() {

	// Set Random Seed
	rand.Seed(time.Now().UnixNano())

	slot := 3
	item := rand.Intn(12)
	luck := rand.Intn(3)

	if e.Items[0].Name == "none" {
		slot = 0
	} else if e.Items[1].Name == "none" {
		slot = 1
	}

	if luck == 1 && slot != 3 {
		switch item {
		case 0, 1:
			e.Items[slot] = itemTemplates["roh"]
		case 2, 3:
			e.Items[slot] = itemTemplates["row"]
		case 4, 5:
			e.Items[slot] = itemTemplates["rob"]
		case 6, 7:
			e.Items[slot] = itemTemplates["rog"]
		case 8:
			e.Items[slot] = itemTemplates["noh"]
		case 9:
			e.Items[slot] = itemTemplates["now"]
		case 10:
			e.Items[slot] = itemTemplates["nob"]
		case 11:
			e.Items[slot] = itemTemplates["nog"]
		}
		// Recalculate Stats with Item
		e.equpItem(slot)
		e.statRecalc()
	} else {
		e.Potions++
	}
}

func (e *Entity) equpItem(slot int) {
	e.Stats[e.Items[slot].Smod] += e.Items[slot].Multi
	e.statRecalc()
}

func (e *Entity) unequpItem(slot int) {
	e.Stats[e.Items[slot].Smod] -= e.Items[slot].Multi
	e.Items[slot] = itemTemplates["none"]
	e.statRecalc()
}

func (e *Entity) usePotion() {
	if e.Potions < 1 {
		fmt.Println("\nNo Potions left... ")
	} else {
		e.Potions--

		php := e.Hp + (50 * e.Lvl) + (e.Stats["sta"] * 10)

		if php > e.MaxHp {
			fmt.Println("Filled HP to the Max!")
			e.Hp = e.MaxHp
		} else {
			fmt.Printf("Healed for %d HP\n", php-e.Hp)
			e.Hp = php
		}

		pmp := e.Mp + (25 * e.Lvl) + (e.Stats["int"] * 10)

		if pmp > e.MaxMp {
			fmt.Println("Filled MP to the Max!")
			e.Mp = e.MaxMp
		} else {
			fmt.Printf("Restored %d MP\n", pmp-e.Mp)
			e.Mp = pmp
		}
	}
}

func (e Entity) showStats() {
	fmt.Print(
		" Name\t", e.Name, "\n",
		" Class\t", e.Class.Name, "\n",
		" Level\t", e.Lvl, "\n")
	if e.Player {
		fmt.Print(" EXP\t", e.Exp, " / ", e.MaxExp, "\n")
	} else {
		fmt.Print(" EXP\t", e.Exp, "\n")
	}
	fmt.Print(
		" \n",
		" Hitpoints\t", e.Hp, " / ", e.MaxHp, "\n",
		" Manapoints\t", e.Mp, " / ", e.MaxMp, "\n",
		" STA ", e.Stats["sta"], "\t\tSTR\t", e.Stats["str"], "\n",
		" DEX ", e.Stats["dex"], "\t\tINT\t", e.Stats["int"], "\n")
	if e.Player {
		fmt.Print(" Itemslot 1:\t", e.Items[0].Name, "\n")
		fmt.Print(" Itemslot 2:\t", e.Items[1].Name, "\n")
	}
	fmt.Print("≡‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗≡\n\n")
}

func (e *Entity) battle(attack int, other *Entity) {

	// Set Random Seed
	rand.Seed(time.Now().UnixNano())

	if e.Mp >= e.Attacks[attack].Cost {
		pdmg := e.Attacks[attack].Damage + rand.Intn(e.Lvl)
		dice := rand.Intn(20) + 1

		if other.Hp > 0 && pdmg < other.Hp {
			if dice == 20 {
				fmt.Println("CRITICAL HIT!", pdmg*2, "DMG!")
				other.Hp -= pdmg * 2
			} else if dice == 1 {
				fmt.Println("Your attacked missed...")
			} else {
				fmt.Println("You hit the", other.Name, "for", pdmg, "dmg")
				other.Hp -= pdmg
			}
		} else {
			other.Hp = 0
			other.Alive = false
		}

		e.Mp -= e.Attacks[attack].Cost

	} else {
		fmt.Println("You don't have enough mana")
	}

	if other.Alive {
		edmg := other.Attacks[1].Damage + rand.Intn(e.Lvl)

		if edmg >= e.Hp {
			e.Hp = 0
			e.Alive = false
			fmt.Println("You died a horrible death...")
		} else {
			if rand.Intn(10) > 8 {
				if other.Mp >= other.Attacks[1].Cost {
					fmt.Println(other.Name, other.Attacks[1].Name, "you for ", edmg, "dmg")
					e.Hp -= edmg
					other.Mp -= other.Attacks[1].Cost
				} else {
					fmt.Println(other.Name, "misses its Attack")
				}
			} else {
				fmt.Println(other.Name, other.Attacks[0].Name, "you for ", edmg, "dmg")
				e.Hp -= edmg
			}
		}
	}
}
