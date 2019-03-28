package main

import (
	"fmt"
	"math/rand"
	"time"
)

func newEnemy(scale int) *Entity {

	// Set Random Seed
	rand.Seed(time.Now().UnixNano())

	// Set Enemy

	e := Entity{}
	var className string
	class := rand.Intn(4)

	switch class {
	case 0:
		className = "humanoid"
		e.Name = "Soldier"
	case 1:
		className = "humanoid"
		e.Name = "Ghoul"
	case 2:
		className = "animal"
		e.Name = "Wolf"
	case 3:
		className = "animal"
		e.Name = "Bat"
	}

	e.class.name = className
	e.class.primary = classTemplates[className].primary
	e.stats = map[string]int{}
	e.mp = 100
	e.maxMp = e.mp
	e.exp = 35 * (scale + scale)
	e.lvl = scale
	e.Player = false
	e.Alive = true
	e.stats["str"] = entityTemplates[className].stats["str"] + scale
	e.stats["sta"] = entityTemplates[className].stats["sta"] + scale
	e.stats["dex"] = entityTemplates[className].stats["dex"] + scale
	e.stats["int"] = entityTemplates[className].stats["int"] + scale
	e.hp = (e.stats["sta"] + rand.Intn(3)) * 10 * scale
	e.maxHp = e.hp

	for _, a := range attackTemplates {
		if a.classreq == e.class.name && e.lvl >= a.lvlreq {
			e.attacks[a.slot] = a
			e.attacks[a.slot].damage = a.damage + e.stats[e.primary]
			e.attacks[a.slot].cost = a.cost + 2*e.lvl
		}
	}

	fmt.Println(e.attacks)

	return &e
}
