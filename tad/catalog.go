package main

var entityTemplates = map[string]entity{
	"mage": entity{
		class: classTemplates["mage"],
		stats: map[string]int{
			"sta": 5,
			"str": 3,
			"dex": 2,
			"int": 10,
		},
	},
	"warrior": entity{
		class: classTemplates["warrior"],
		stats: map[string]int{
			"sta": 7,
			"str": 7,
			"dex": 4,
			"int": 2,
		},
	},
	"paladin": entity{
		class: classTemplates["paladin"],
		stats: map[string]int{
			"sta": 5,
			"str": 6,
			"dex": 4,
			"int": 5,
		},
	},
	"ranger": entity{
		class: classTemplates["ranger"],
		stats: map[string]int{
			"sta": 6,
			"str": 2,
			"dex": 8,
			"int": 4,
		},
	},
	"humanoid": entity{
		class: classTemplates["humanoid"],
		stats: map[string]int{
			"sta": 5,
			"str": 5,
			"dex": 5,
			"int": 5,
		},
	},
	"animal": entity{
		class: classTemplates["animal"],
		stats: map[string]int{
			"sta": 6,
			"str": 4,
			"dex": 6,
			"int": 4,
		},
	},
}

var classTemplates = map[string]class{
	"mage": class{
		name:    "Mage",
		primary: "int",
	},
	"warrior": class{
		name:    "Warrior",
		primary: "str",
	},
	"paladin": class{
		name:    "Paladin",
		primary: "sta",
	},
	"ranger": class{
		name:    "Ranger",
		primary: "dex",
	},
	"any": class{
		name:    "any",
		primary: "str",
	},
	"humanoid": class{
		name:    "Humanoid",
		primary: "str",
	},
	"animal": class{
		name:    "Animal",
		primary: "dex",
	},
}

var attackTemplates = map[string]attack{

	"Unarmed Strike": attack{
		name:     "Unarmed Strike",
		damage:   25,
		lvlreq:   1,
		classreq: classTemplates["any"],
		slot:     0,
	},

	//// Player Atks
	// WARRIOR
	"Obliterate": attack{
		name:     "Obliterate",
		damage:   30,
		lvlreq:   2,
		classreq: classTemplates["warrior"],
		cost:     5,
		slot:     1,
	},
	"Mortal Strike": attack{
		name:     "Mortal Strike",
		damage:   40,
		lvlreq:   3,
		classreq: classTemplates["warrior"],
		cost:     15,
		slot:     2,
	},
	"Execute": attack{
		name:     "Execute",
		damage:   80,
		lvlreq:   5,
		classreq: classTemplates["warrior"],
		cost:     50,
		slot:     3,
	},
	// MAGE
	"Lightning Bolt": attack{
		name:     "Lightning Bolt",
		damage:   25,
		lvlreq:   2,
		classreq: classTemplates["mage"],
		cost:     5,
		slot:     1,
	},
	"Fireball": attack{
		name:     "Fireball",
		damage:   35,
		lvlreq:   3,
		classreq: classTemplates["mage"],
		cost:     8,
		slot:     2,
	},
	"Ice Beam": attack{
		name:     "Ice Beam",
		damage:   75,
		lvlreq:   5,
		classreq: classTemplates["mage"],
		cost:     50,
		slot:     3,
	},
	// RANGER
	"Cobra Shot": attack{
		name:     "Cobra Shot",
		damage:   25,
		lvlreq:   2,
		classreq: classTemplates["ranger"],
		cost:     5,
		slot:     1,
	},
	"Poison Arrow": attack{
		name:     "Poison Arrow",
		damage:   35,
		lvlreq:   3,
		classreq: classTemplates["ranger"],
		cost:     10,
		slot:     2,
	},
	"Rain of Arrows": attack{
		name:     "Rain of Arrows",
		damage:   70,
		lvlreq:   5,
		classreq: classTemplates["ranger"],
		cost:     45,
		slot:     3,
	},
	// PALADIN
	"Holy Light": attack{
		name:     "Holy Light",
		damage:   25,
		lvlreq:   2,
		classreq: classTemplates["paladin"],
		cost:     5,
		slot:     1,
	},
	"Divine Storm": attack{
		name:     "Divine Storm",
		damage:   40,
		lvlreq:   3,
		classreq: classTemplates["paladin"],
		cost:     12,
		slot:     2,
	},
	"Sword of Justice": attack{
		name:     "Sword of Justice",
		damage:   65,
		lvlreq:   5,
		classreq: classTemplates["paladin"],
		cost:     45,
		slot:     3,
	},
	//// Enemy Attacks
	// HUMANOIDS
	"Punch": attack{
		name:     "punches",
		damage:   3,
		lvlreq:   1,
		classreq: classTemplates["humanoid"],
		cost:     1,
		slot:     0,
	},
	"Strike": attack{
		name:     "strikes",
		damage:   6,
		lvlreq:   1,
		classreq: classTemplates["humanoid"],
		cost:     25,
		slot:     1,
	},
	// ANIMALS
	"Bite": attack{
		name:     "bites",
		damage:   3,
		lvlreq:   1,
		classreq: classTemplates["animal"],
		cost:     1,
		slot:     0,
	},
	"Scratche": attack{
		name:     "scratches",
		damage:   6,
		lvlreq:   1,
		classreq: classTemplates["animal"],
		cost:     25,
		slot:     1,
	},
}
