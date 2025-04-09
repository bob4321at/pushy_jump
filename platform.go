package main

import (
	"image/color"

	"github.com/gen2brain/raylib-go/raylib"
)

type Platform struct {
	Pos   rl.Vector3
	Size  rl.Vector3
	Color color.Color
}

func NewPlatform(pos, size rl.Vector3, color color.Color) (p Platform) {
	p.Pos = pos
	p.Size = size
	p.Color = color

	return p
}

func (p *Platform) Draw() {
	r, g, b, a := p.Color.RGBA()
	rl.DrawModelEx(cube_mesh, p.Pos, rl.NewVector3(0, 0, 0), 0, p.Size, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
}

var platforms = []Platform{
	NewPlatform(rl.NewVector3(0, -5, 0), rl.NewVector3(100, 1, 100), color.RGBA{100, 100, 255, 255}),
	NewPlatform(rl.NewVector3(20, -5, 0), rl.NewVector3(10, 10, 10), color.RGBA{100, 100, 255, 255}),
	NewPlatform(rl.NewVector3(20, 25, 0), rl.NewVector3(10, 10, 10), color.RGBA{100, 100, 255, 255}),
	NewPlatform(rl.NewVector3(-20, -3, 0), rl.NewVector3(1, 1, 1), color.RGBA{100, 100, 255, 255}),
}
