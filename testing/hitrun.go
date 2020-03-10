package main

import tl "github.com/JoelOtter/termloop"

type Character struct {
	*tl.Rectangle
	move  bool
	npc   bool
	enemy bool
	px    int
	py    int
}

func NewCharacter(x, y, w, h int, color tl.Attr, move bool, npc bool, enemy bool) *Character {
	return &Character{
		Rectangle: tl.NewRectangle(x, y, w, h, color),
		move:      move,
		npc:       npc,
		enemy:     enemy,
	}
}

func (c *Character) Tick(ev tl.Event) {
	if ev.Type == tl.EventKey && c.move {
		c.px, c.py = c.Position()
		switch ev.Key {
		case tl.KeyArrowRight:
			c.SetPosition(c.px+1, c.py)
		case tl.KeyArrowLeft:
			c.SetPosition(c.px-1, c.py)
		case tl.KeyArrowDown:
			c.SetPosition(c.px, c.py+1)
		case tl.KeyArrowUp:
			c.SetPosition(c.px, c.py-1)
		case tl.KeySpace:
			c.SetColor(tl.ColorYellow)
		}
	}
}

func (c *Character) Collide(p tl.Physical) {
	if _, ok := p.(*Character); ok && c.move {
		c.SetPosition(c.px, c.py)
	}
	if _, ok := p.(*Character); ok && c.enemy {
		c.SetColor(tl.ColorRed)
		c.SetPosition(c.px, c.py)
	}
}

func main() {
	game := tl.NewGame()
	game.Screen().SetFps(60)

	lvl := tl.NewBaseLevel(tl.Cell{Bg: tl.ColorWhite})
	lvl.AddEntity(NewCharacter(2, 2, 1, 1, tl.ColorYellow, true, false, false))
	lvl.AddEntity(NewCharacter(10, 9, 1, 1, tl.ColorRed, false, true, true))
	lvl.AddEntity(NewCharacter(8, 4, 1, 1, tl.ColorGreen, false, true, false))
	lvl.AddEntity(NewCharacter(8, 6, 3, 2, tl.ColorBlack, false, false, false))
	lvl.AddEntity(NewCharacter(12, 7, 3, 2, tl.ColorBlack, false, false, false))
	lvl.AddEntity(NewCharacter(14, 9, 3, 2, tl.ColorBlack, false, false, false))

	game.Screen().SetLevel(lvl)
	game.Screen().AddEntity(tl.NewFpsText(1, 1, tl.ColorRed, tl.ColorDefault, 0.5))
	game.Start()
}
