package main

var entityTemplates = map[string]Entity{
	"mage": Entity{
		class: classTemplates["mage"],
		stats: map[string]int{
			"sta": 5,
			"str": 3,
			"dex": 2,
			"int": 10,
		},
	},
	"warrior": Entity{
		class: classTemplates["warrior"],
		stats: map[string]int{
			"sta": 7,
			"str": 7,
			"dex": 4,
			"int": 2,
		},
	},
	"paladin": Entity{
		class: classTemplates["paladin"],
		stats: map[string]int{
			"sta": 5,
			"str": 6,
			"dex": 4,
			"int": 5,
		},
	},
	"ranger": Entity{
		class: classTemplates["ranger"],
		stats: map[string]int{
			"sta": 6,
			"str": 2,
			"dex": 8,
			"int": 4,
		},
	},
	"humanoid": Entity{
		class: classTemplates["humanoid"],
		stats: map[string]int{
			"sta": 4,
			"str": 5,
			"dex": 6,
			"int": 5,
		},
	},
	"animal": Entity{
		class: classTemplates["animal"],
		stats: map[string]int{
			"sta": 4,
			"str": 6,
			"dex": 4,
			"int": 6,
		},
	},
}

var classTemplates = map[string]class{
	"mage": class{
		name:    "mage",
		primary: "int",
	},
	"warrior": class{
		name:    "warrior",
		primary: "str",
	},
	"paladin": class{
		name:    "paladin",
		primary: "sta",
	},
	"ranger": class{
		name:    "ranger",
		primary: "dex",
	},
	"any": class{
		name:    "any",
		primary: "str",
	},
	"humanoid": class{
		name:    "humanoid",
		primary: "str",
	},
	"animal": class{
		name:    "animal",
		primary: "dex",
	},
}

var attackTemplates = map[string]attack{

	"Unarmed Strike": attack{
		name:     "Unarmed Strike",
		damage:   25,
		lvlreq:   1,
		classreq: "any",
		slot:     0,
	},

	//// Player Atks
	// WARRIOR
	"Obliterate": attack{
		name:     "Obliterate",
		damage:   30,
		lvlreq:   2,
		classreq: "warrior",
		cost:     5,
		slot:     1,
	},
	"Mortal Strike": attack{
		name:     "Mortal Strike",
		damage:   40,
		lvlreq:   3,
		classreq: "warrior",
		cost:     15,
		slot:     2,
	},
	"Execute": attack{
		name:     "Execute",
		damage:   100,
		lvlreq:   5,
		classreq: "warrior",
		cost:     50,
		slot:     3,
	},
	// MAGE
	"Lightning Bolt": attack{
		name:     "Lightning Bolt",
		damage:   25,
		lvlreq:   2,
		classreq: "mage",
		cost:     5,
		slot:     1,
	},
	"Fireball": attack{
		name:     "Fireball",
		damage:   35,
		lvlreq:   3,
		classreq: "mage",
		cost:     8,
		slot:     2,
	},
	"Ice Beam": attack{
		name:     "Ice Beam",
		damage:   95,
		lvlreq:   5,
		classreq: "mage",
		cost:     50,
		slot:     3,
	},
	// RANGER
	"Cobra Shot": attack{
		name:     "Cobra Shot",
		damage:   25,
		lvlreq:   2,
		classreq: "ranger",
		cost:     5,
		slot:     1,
	},
	"Poison Arrow": attack{
		name:     "Poison Arrow",
		damage:   35,
		lvlreq:   3,
		classreq: "ranger",
		cost:     10,
		slot:     2,
	},
	"Rain of Arrows": attack{
		name:     "Rain of Arrows",
		damage:   90,
		lvlreq:   5,
		classreq: "ranger",
		cost:     45,
		slot:     3,
	},
	// PALADIN
	"Holy Light": attack{
		name:     "Holy Light",
		damage:   25,
		lvlreq:   2,
		classreq: "paladin",
		cost:     5,
		slot:     1,
	},
	"Divine Storm": attack{
		name:     "Divine Storm",
		damage:   40,
		lvlreq:   3,
		classreq: "paladin",
		cost:     12,
		slot:     2,
	},
	"Sword of Justice": attack{
		name:     "Sword of Justice",
		damage:   95,
		lvlreq:   5,
		classreq: "paladin",
		cost:     45,
		slot:     3,
	},
	//// Enemy Attacks
	// HUMANOIDS
	"Punch": attack{
		name:     "punches",
		damage:   3,
		lvlreq:   1,
		classreq: "humanoid",
		cost:     1,
		slot:     0,
	},
	"Strike": attack{
		name:     "strikes",
		damage:   6,
		lvlreq:   1,
		classreq: "humanoid",
		cost:     25,
		slot:     1,
	},
	// ANIMALS
	"Bite": attack{
		name:     "bites",
		damage:   3,
		lvlreq:   1,
		classreq: "animal",
		cost:     1,
		slot:     0,
	},
	"Scratche": attack{
		name:     "scratches",
		damage:   6,
		lvlreq:   1,
		classreq: "animal",
		cost:     25,
		slot:     1,
	},
}

var itemTemplates = map[string]item{
	"row": item{
		name:  "Ring of Wisdom",
		multi: 5,
		smod:  "int",
	},
	"roh": item{
		name:  "Ring of Haste",
		multi: 5,
		smod:  "dex",
	},
	"rog": item{
		name:  "Ring of Giants",
		multi: 5,
		smod:  "sta",
	},
	"rob": item{
		name:  "Ring of the Berserker",
		multi: 5,
		smod:  "str",
	},
	"none": item{
		name:  "none",
		multi: 0,
		smod:  "str",
	},
}
