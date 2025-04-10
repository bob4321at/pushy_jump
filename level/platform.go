package platform

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	Cube_Mesh rl.Model
)

type Platform struct {
	Pos  rl.Vector3
	Size rl.Vector3
}

func NewPlatform(pos, size rl.Vector3) (p Platform) {
	p.Pos = pos
	p.Size = size

	return p
}

func (p *Platform) Draw() {
	rl.DrawModelEx(Cube_Mesh, p.Pos, rl.NewVector3(0, 0, 0), 0, p.Size, rl.White)
}

var Platforms = []Platform{}
