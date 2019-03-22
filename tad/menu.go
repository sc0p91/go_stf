package main

import (
	"fmt"
)

func menu() {
	clear()
	fmt.Print(
		"‗========== PLAYER ===========‗\n",
		"  Name\t", pname, "\n",
		"  Class\t", pclass, "\n",
		"  Level\t", plvl, "\n",
		"  EXP\t", pexp, " / ", mexp, " \n",
		"  \n",
		"  Hitpoints\t", hp, " / ", max_hp, "\n",
		"  Manapoints\t", mp, " / ", max_mp, "\n",
		"  STA ", psta, "\t\tSTR\t", pstr, "\n",
		"  DEX ", pdex, "\t\tINT\t", pint, "\n",
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