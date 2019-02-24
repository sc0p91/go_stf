package main

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Demo",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	sqr := imdraw.New(nil)
	tri := imdraw.New(nil)
	sgn := imdraw.New(nil)

	sqr.Color = pixel.RGB(0, 1, 0)
	sqr.Push(pixel.V(0, 0))
	sqr.Push(pixel.V(1024, 0))
	sqr.Color = pixel.RGB(0, 0, 1)
	sqr.Push(pixel.V(1024, 1024))
	sqr.Push(pixel.V(0, 1024))
	sqr.Polygon(0)

	tri.Color = pixel.RGB(1, 0.5, 0)
	tri.Push(pixel.V(800, 50))
	tri.Push(pixel.V(1000, 50))
	tri.Color = pixel.RGB(1, 0.5, 0.5)
	tri.Push(pixel.V(900, 200))
	tri.Polygon(0)

	sgn.Color = pixel.RGB(0.3, 0.6, 0.9)
	sgn.EndShape = imdraw.RoundEndShape
	sgn.Push(pixel.V(800, 500), pixel.V(750, 50))
	sgn.Push(pixel.V(750, 50), pixel.V(50, 50))
	sgn.EndShape = imdraw.SharpEndShape
	sgn.Push(pixel.V(750, 720), pixel.V(50, 720))
	sgn.Push(pixel.V(50, 720), pixel.V(100, 500))
	sgn.Line(30)
	sgn.Color = pixel.RGB(0.4, 0.7, 0.1)
	sgn.Push(pixel.V(200, 500), pixel.V(800, 500))
	sgn.Ellipse(pixel.V(120, 80), 0)
	sgn.Color = pixel.RGB(0.6, 0.3, 0)
	sgn.Push(pixel.V(500, 350))
	sgn.Circle(100, 30)
	sgn.Color = pixel.RGB(0.1, 0.2, 0.9)
	sgn.Push(pixel.V(500, 350))
	sgn.CircleArc(150, -math.Pi, 0, 30)

	for !win.Closed() {

		win.Clear(colornames.Gray)

		sqr.Draw(win)
		tri.Draw(win)
		sgn.Draw(win)

		win.Update()

	}
}

func main() {
	pixelgl.Run(run)
}
