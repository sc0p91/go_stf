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
		stats: map[string]int {
			"sta": 7,
			"str": 7,
			"dex": 4,
			"int": 2,
		},
	},
	"paladin": entity{
		class: classTemplates["paladin"],
		stats: map[string]int {
			"sta": 5,
			"str": 6,
			"dex": 4,
			"int": 5,
		},
	},
	"ranger": entity{
		class: classTemplates["ranger"],
		stats: map[string]int {
			"sta": 6,
			"str": 2,
			"dex": 8,
			"int": 4,
		},
	},
}

var classTemplates = map[string]class {
	"mage": class{
		name: "Mage",
		primary: "int",
	},
	"warrior": class{
		name: "Warrior",
		primary: "str",
	},
	"paladin": class{
		name: "Paladin",
		primary: "sta",
	},
	"ranger": class{
		name: "Ranger",
		primary: "dex",
	},
	"any": class{
		name: "any",
		primary: "str",
	},
	"humanoid": class{
		name: "Humanoid",
		primary: "str",
	},
	"animal": class{
		name: "Animal",
		primary: "dex",
	},
}

var attackTemplates = map[string]attack {
	"Unarmed Strike": attack{
		name: "Unarmed Strike",
		damage: 25,
		lvlreq: 1,
		classreq: classTemplates["any"],
		slot: 0,
	},
	"Obliterate": attack{
		name: "Obliterate",
		damage: 40,
		lvlreq: 2,
		classreq: classTemplates["warrior"],
		cost: 5,
		slot: 1,
	},
	"Lightning Bolt": attack{
		name: "Lightning Bolt",
		damage: 25,
		lvlreq: 2,
		classreq: classTemplates["mage"],
		cost: 5,
		slot: 1,
	},
	"Fireball": attack{
		name: "Fireball",
		damage: 50,
		lvlreq: 3,
		classreq: classTemplates["mage"],
		cost: 8,
		slot: 2,
	},
	"Cobra Shot": attack{
		name: "Cobra Shot",
		damage: 25,
		lvlreq: 2,
		classreq: classTemplates["ranger"],
		cost: 5,
		slot: 1,
	},
}