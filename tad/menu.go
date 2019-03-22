package main

import (
	"fmt"
)

func (p *player) menu() {
	clear()
	fmt.Print(
		"‗========== PLAYER ===========‗\n",
		"  Name\t", p.name, "\n",
		"  Class\t", p.class, "\n",
		"  Level\t", p.lvl, "\n",
		"  EXP\t", p.exp, " / ", p.maxExp, " \n",
		"  \n",
		"  Hitpoints\t", p.hp, " / ", p.maxHp, "\n",
		"  Manapoints\t", p.mp, " / ", p.maxMp, "\n",
		"  STA ", p.sta, "\t\tSTR\t", p.str, "\n",
		"  DEX ", p.dex, "\t\tINT\t", p.int, "\n",
		"≡‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗≡\n",
		"\n",
		"‗========= ACTIONS ==========‗\n",
		"  [Q] Placeholder Atk1 \n",
		"  [W] Placeholder Atk2 \n",
		"  [E] Placeholder Atk3 \n",
		"  [R] Placeholder Atk4 \n",
		"  [A] Placeholder Item1\n",
		"  [S] Placeholder Item2\n",
		"  [D] Restore HP|MP    \n",
		"  [F] Do nothing.      \n",
		"  [X] quit             \n",
		"≡‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗‗≡\n")
}