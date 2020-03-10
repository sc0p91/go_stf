package main

import tl "github.com/JoelOtter/termloop"

type Character struct {
	*tl.Rectangle
	move bool
	px   int
	py   int
}

type NPC struct {
	*tl.Rectangle
	move  bool
	enemy bool
	px    int
	py    int
}

func NewNPC(x, y, w, h int, color tl.Attr, move bool, enemy bool) *NPC {
	return &NPC{
		Rectangle: tl.NewRectangle(x, y, w, h, color),
		move:      move,
		enemy:     enemy,
	}
}

func NewCharacter(x, y, w, h int, color tl.Attr, move bool) *Character {
	return &Character{
		Rectangle: tl.NewRectangle(x, y, w, h, color),
		move:      move,
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
	//if _, ok := p.(*NPC); ok && c.move {
	if _, ok := p.(*NPC); ok {
		c.SetPosition(c.px, c.py)
	}
}

func main() {
	game := tl.NewGame()
	game.Screen().SetFps(60)

	lvl := tl.NewBaseLevel(tl.Cell{Bg: tl.ColorWhite})
	lvl.AddEntity(NewCharacter(2, 2, 1, 1, tl.ColorYellow, true))
	lvl.AddEntity(NewNPC(10, 9, 1, 1, tl.ColorRed, false, true))
	lvl.AddEntity(NewNPC(8, 4, 1, 1, tl.ColorGreen, false, true))
	lvl.AddEntity(NewNPC(8, 6, 3, 2, tl.ColorBlack, false, false))
	lvl.AddEntity(NewNPC(12, 7, 3, 2, tl.ColorBlack, false, false))
	lvl.AddEntity(NewNPC(14, 9, 3, 2, tl.ColorBlack, false, false))
	lvl.AddEntity(NewNPC(0, 0, 1, 20, tl.ColorBlack, false, false))
	lvl.AddEntity(NewNPC(0, 0, 60, 1, tl.ColorBlack, false, false))
	lvl.AddEntity(NewNPC(0, 20, 60, 1, tl.ColorBlack, false, false))
	lvl.AddEntity(NewNPC(60, 0, 1, 21, tl.ColorBlack, false, false))

	game.Screen().SetLevel(lvl)
	game.Screen().AddEntity(tl.NewFpsText(0, 22, tl.ColorRed, tl.ColorDefault, 0.5))
	game.Start()
}
