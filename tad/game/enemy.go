package game

import (
	"math/rand"
	"time"
)

func NewEnemy(scale int) *Entity {

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

	e.Class.Name = className
	e.Class.Primary = classTemplates[className].Primary
	e.Stats = map[string]int{}
	e.Mp = 100
	e.MaxMp = e.Mp
	e.Exp = 35 * (scale + scale)
	e.Lvl = scale
	e.Player = false
	e.Alive = true
	e.Stats["str"] = entityTemplates[className].Stats["str"] + scale
	e.Stats["sta"] = entityTemplates[className].Stats["sta"] + scale
	e.Stats["dex"] = entityTemplates[className].Stats["dex"] + scale
	e.Stats["int"] = entityTemplates[className].Stats["int"] + scale
	e.Hp = (e.Stats["sta"] + rand.Intn(3)) * 10 * scale
	e.MaxHp = e.Hp

	for _, a := range attackTemplates {
		if a.ClassReq == e.Class.Name && e.Lvl >= a.LvlReq {
			e.Attacks[a.Slot] = a
			e.Attacks[a.Slot].Damage = a.Damage + e.Stats[e.Primary]
			e.Attacks[a.Slot].Cost = a.Cost + 2*e.Lvl
		}
	}

	return &e
}
