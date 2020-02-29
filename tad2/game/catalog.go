package game

var entityTemplates = map[string]Entity{
	"mage": Entity{
		Class: classTemplates["mage"],
		Stats: map[string]int{
			"sta": 5,
			"str": 3,
			"dex": 2,
			"int": 10,
		},
	},
	"warrior": Entity{
		Class: classTemplates["warrior"],
		Stats: map[string]int{
			"sta": 7,
			"str": 7,
			"dex": 4,
			"int": 2,
		},
	},
	"paladin": Entity{
		Class: classTemplates["paladin"],
		Stats: map[string]int{
			"sta": 5,
			"str": 6,
			"dex": 4,
			"int": 5,
		},
	},
	"ranger": Entity{
		Class: classTemplates["ranger"],
		Stats: map[string]int{
			"sta": 6,
			"str": 2,
			"dex": 8,
			"int": 4,
		},
	},
	"humanoid": Entity{
		Class: classTemplates["humanoid"],
		Stats: map[string]int{
			"sta": 4,
			"str": 5,
			"dex": 6,
			"int": 5,
		},
	},
	"animal": Entity{
		Class: classTemplates["animal"],
		Stats: map[string]int{
			"sta": 4,
			"str": 6,
			"dex": 4,
			"int": 6,
		},
	},
}

var classTemplates = map[string]Class{
	"mage": Class{
		Name:    "mage",
		Primary: "int",
	},
	"warrior": Class{
		Name:    "warrior",
		Primary: "str",
	},
	"paladin": Class{
		Name:    "paladin",
		Primary: "sta",
	},
	"ranger": Class{
		Name:    "ranger",
		Primary: "dex",
	},
	"any": Class{
		Name:    "any",
		Primary: "str",
	},
	"humanoid": Class{
		Name:    "humanoid",
		Primary: "str",
	},
	"animal": Class{
		Name:    "animal",
		Primary: "dex",
	},
}

var attackTemplates = map[string]Attack{

	"Unarmed Strike": Attack{
		Name:     "Unarmed Strike",
		Damage:   25,
		LvlReq:   1,
		ClassReq: "any",
		Slot:     0,
	},

	//// Player Atks
	// WARRIOR
	"Obliterate": Attack{
		Name:     "Obliterate",
		Damage:   30,
		LvlReq:   2,
		ClassReq: "warrior",
		Cost:     5,
		Slot:     1,
	},
	"Mortal Strike": Attack{
		Name:     "Mortal Strike",
		Damage:   40,
		LvlReq:   3,
		ClassReq: "warrior",
		Cost:     15,
		Slot:     2,
	},
	"Execute": Attack{
		Name:     "Execute",
		Damage:   100,
		LvlReq:   5,
		ClassReq: "warrior",
		Cost:     50,
		Slot:     3,
	},
	// MAGE
	"Lightning Bolt": Attack{
		Name:     "Lightning Bolt",
		Damage:   25,
		LvlReq:   2,
		ClassReq: "mage",
		Cost:     5,
		Slot:     1,
	},
	"Fireball": Attack{
		Name:     "Fireball",
		Damage:   35,
		LvlReq:   3,
		ClassReq: "mage",
		Cost:     8,
		Slot:     2,
	},
	"Ice Beam": Attack{
		Name:     "Ice Beam",
		Damage:   95,
		LvlReq:   5,
		ClassReq: "mage",
		Cost:     50,
		Slot:     3,
	},
	// RANGER
	"Cobra Shot": Attack{
		Name:     "Cobra Shot",
		Damage:   25,
		LvlReq:   2,
		ClassReq: "ranger",
		Cost:     5,
		Slot:     1,
	},
	"Poison Arrow": Attack{
		Name:     "Poison Arrow",
		Damage:   35,
		LvlReq:   3,
		ClassReq: "ranger",
		Cost:     10,
		Slot:     2,
	},
	"Rain of Arrows": Attack{
		Name:     "Rain of Arrows",
		Damage:   90,
		LvlReq:   5,
		ClassReq: "ranger",
		Cost:     45,
		Slot:     3,
	},
	// PALADIN
	"Holy Light": Attack{
		Name:     "Holy Light",
		Damage:   25,
		LvlReq:   2,
		ClassReq: "paladin",
		Cost:     5,
		Slot:     1,
	},
	"Divine Storm": Attack{
		Name:     "Divine Storm",
		Damage:   40,
		LvlReq:   3,
		ClassReq: "paladin",
		Cost:     12,
		Slot:     2,
	},
	"Sword of Justice": Attack{
		Name:     "Sword of Justice",
		Damage:   95,
		LvlReq:   5,
		ClassReq: "paladin",
		Cost:     45,
		Slot:     3,
	},
	//// Enemy Attacks
	// HUMANOIDS
	"Punch": Attack{
		Name:     "punches",
		Damage:   3,
		LvlReq:   1,
		ClassReq: "humanoid",
		Cost:     1,
		Slot:     0,
	},
	"Strike": Attack{
		Name:     "strikes",
		Damage:   6,
		LvlReq:   1,
		ClassReq: "humanoid",
		Cost:     25,
		Slot:     1,
	},
	// ANIMALS
	"Bite": Attack{
		Name:     "bites",
		Damage:   3,
		LvlReq:   1,
		ClassReq: "animal",
		Cost:     1,
		Slot:     0,
	},
	"Scratche": Attack{
		Name:     "scratches",
		Damage:   6,
		LvlReq:   1,
		ClassReq: "animal",
		Cost:     25,
		Slot:     1,
	},
}

var itemTemplates = map[string]Item{
	"row": Item{
		Name:  "Ring of Wisdom",
		Multi: 5,
		Smod:  "int",
	},
	"roh": Item{
		Name:  "Ring of Haste",
		Multi: 5,
		Smod:  "dex",
	},
	"rog": Item{
		Name:  "Ring of Giants",
		Multi: 5,
		Smod:  "sta",
	},
	"rob": Item{
		Name:  "Ring of the Berserker",
		Multi: 5,
		Smod:  "str",
	},
	"now": Item{
		Name:  "Necklace of Wisdom",
		Multi: 5,
		Smod:  "int",
	},
	"noh": Item{
		Name:  "Necklace of Haste",
		Multi: 10,
		Smod:  "dex",
	},
	"nog": Item{
		Name:  "Necklace of Giants",
		Multi: 10,
		Smod:  "sta",
	},
	"nob": Item{
		Name:  "Necklace of the Berserker",
		Multi: 10,
		Smod:  "str",
	},
	"none": Item{
		Name:  "none",
		Multi: 0,
		Smod:  "str",
	},
}
