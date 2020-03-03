package main

import tl "github.com/JoelOtter/termloop"

//Player creation
type Player struct {
	*tl.Entity
	prevX int
	prevY int
	level *tl.BaseLevel
}

//Physical For static stuff ie. Walls
type Physical interface {
	Position() (int, int)
	Size() (int, int)
}

//DynamicPhysical for dynamic stuff ie. Player
type DynamicPhysical interface {
	Position() (int, int)
	Size() (int, int)
	Collide(Physical)
}

//Draw to move the cam around player
func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.Position()
	player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	player.Entity.Draw(screen)
}

//Tick set for movement
func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		player.prevX, player.prevY = player.Position()
		switch event.Key {
		case tl.KeyArrowRight:
			player.SetPosition(player.prevX+1, player.prevY)
		case tl.KeyArrowLeft:
			player.SetPosition(player.prevX-1, player.prevY)
		case tl.KeyArrowUp:
			player.SetPosition(player.prevX, player.prevY-1)
		case tl.KeyArrowDown:
			player.SetPosition(player.prevX, player.prevY+1)
		}
	}
}

//Size to return player size
func (player *Player) Size() (int, int) {
	return player.Size()
}

//Position to return player position
func (player *Player) Position() (int, int) {
	return player.Position()
}

//Collide to check player colliding with env
func (player *Player) Collide(collision tl.Physical) {
	if _, ok := collision.(*tl.Rectangle); ok {
		player.SetPosition(player.prevX, player.prevY)
	}
}

//main fuction
func main() {

	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: '.',
	})

	player := Player{
		Entity: tl.NewEntity(1, 1, 1, 1),
		level:  level,
	}
	player.SetCell(0, 0, &tl.Cell{
		Fg: tl.ColorBlack,
		Ch: '8',
	})

	level.AddEntity(&player)
	level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorCyan))

	game.Screen().SetLevel(level)
	game.Start()

}
