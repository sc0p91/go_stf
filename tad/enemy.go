package main

import (
	"math/rand"
	"time"
)

func newEnemy(scale int) *entity {

	// Set Random Seed
	rand.Seed(time.Now().UnixNano())

	// Set Enemy
	e := entity{}

	class := rand.Intn(4)

	switch class {
	case 0:
		e = entityTemplates["humanoid"]
		e.name = "Soldier"
	case 1:
		e = entityTemplates["humanoid"]
		e.name = "Ghoul"
	case 2:
		e = entityTemplates["animal"]
		e.name = "Wolf"
	case 3:
		e = entityTemplates["animal"]
		e.name = "Bat"
	}

	e.hp = (e.stats["sta"] + rand.Intn(5)) * 10 * scale
	e.mp = 100
	e.maxHp = e.hp
	e.maxMp = e.mp
	e.exp = 30 * scale
	e.lvl = scale
	e.player = false
	e.alive = true
	e.stats["str"] += scale
	e.stats["sta"] += scale
	e.stats["dex"] += scale
	e.stats["int"] += scale

	for _, a := range attackTemplates {
		if a.classreq.name == e.class.name && e.lvl >= a.lvlreq {
			e.attacks[a.slot] = a
			e.attacks[a.slot].damage = a.damage + e.stats[e.primary]
			e.attacks[a.slot].cost = a.cost + 2*e.lvl
		}
	}

	return &e
}
