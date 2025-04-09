package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
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
	rl.DrawModelEx(cube_mesh, p.Pos, rl.NewVector3(0, 0, 0), 0, p.Size, rl.White)
}

var platforms = []Platform{}
