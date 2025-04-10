package player

import (
	"fmt"
	"math"
	"pushy/camera"
	platform "pushy/level"
	"pushy/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerStruct struct {
	Pos        rl.Vector3
	Vel        rl.Vector3
	Fake_Y_Vel float32
	Speed      float64
}

func NewPlayer(pos rl.Vector3) (p PlayerStruct) {
	p.Pos = pos
	p.Vel = rl.NewVector3(0, 0, 0)
	p.Speed = 5

	return p
}

func (p *PlayerStruct) Update() {
	p.Movement()

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		hit_pos, hit, ray_length := utils.Raycast(rl.Vector3Add(p.Pos, rl.NewVector3(0, 1.5, 0)), camera.Camera.Rot, 100)
		fmt.Println(ray_length)
		if hit {
			platform.Platforms = append(platform.Platforms, platform.NewPlatform(hit_pos, rl.NewVector3(1, 1, 1)))
		}
	}

	camera.Camera.Pos = rl.Vector3Add(p.Pos, rl.NewVector3(0, 1.5, 0))
}

func (p *PlayerStruct) Movement() {
	move_dir := rl.NewVector3(0, 0, 0)
	move_dir.X = float32(math.Cos(utils.Deg2Rad(float64(camera.Camera.Rot.Y)))) * rl.GetFrameTime() * float32(p.Speed)
	move_dir.Z = float32(math.Sin(utils.Deg2Rad(float64(camera.Camera.Rot.Y)))) * rl.GetFrameTime() * float32(p.Speed)
	move_dir_side := rl.NewVector3(0, 0, 0)
	move_dir_side.X = float32(math.Cos(utils.Deg2Rad(float64(camera.Camera.Rot.Y+90)))) * rl.GetFrameTime() * float32(p.Speed)
	move_dir_side.Z = float32(math.Sin(utils.Deg2Rad(float64(camera.Camera.Rot.Y+90)))) * rl.GetFrameTime() * float32(p.Speed)

	// p.Vel = rl.NewVector3(0, p.Vel.Y, 0)

	p.Fake_Y_Vel -= 100 * rl.GetFrameTime() // + (3 * p.Fake_Y_Vel)

	key_hit := 0

	if rl.IsKeyDown(rl.KeyW) {
		p.Vel = rl.Vector3Add(rl.NewVector3(move_dir.X, move_dir.Y, move_dir.Z), p.Vel)
	} else if rl.IsKeyDown(rl.KeyS) {
		p.Vel = rl.Vector3Add(rl.NewVector3(-move_dir.X, -move_dir.Y, -move_dir.Z), p.Vel)
	}

	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyS) {
		key_hit += 1
	}

	if rl.IsKeyDown(rl.KeyD) {
		p.Vel = rl.Vector3Add(rl.NewVector3(move_dir_side.X, move_dir_side.Y, move_dir_side.Z), p.Vel)
	} else if rl.IsKeyDown(rl.KeyA) {
		p.Vel = rl.Vector3Add(rl.NewVector3(-move_dir_side.X, -move_dir_side.Y, -move_dir_side.Z), p.Vel)
	}

	p.Vel = rl.Vector3Add(p.Vel, rl.NewVector3(p.Vel.X/-8, 0, p.Vel.Z/-8))

	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyD) {
		key_hit += 1
	}

	if key_hit >= 2 {
		p.Vel = rl.NewVector3(p.Vel.X/1.1, p.Vel.Y, p.Vel.Z/1.1)
	}
	p.Vel = rl.NewVector3(p.Vel.X, p.Vel.Y, p.Vel.Z)

	for pi := 0; pi < len(platform.Platforms); pi++ {
		platform := &platform.Platforms[pi]
		if utils.Collide(rl.NewVector3(p.Pos.X, p.Pos.Y+(p.Fake_Y_Vel*rl.GetFrameTime()), p.Pos.Z), rl.NewVector3(1, 2, 1), platform.Pos, platform.Size) {
			falling := false
			if p.Fake_Y_Vel < 0 {
				falling = true
			}
			p.Fake_Y_Vel = 0
			if rl.IsKeyDown(rl.KeySpace) && falling {
				p.Fake_Y_Vel = 50
			}
		}
		if utils.Collide(rl.NewVector3(p.Pos.X+p.Vel.X, p.Pos.Y, p.Pos.Z), rl.NewVector3(1, 2, 1), platform.Pos, platform.Size) {
			p.Vel.X = 0
		}
		if utils.Collide(rl.NewVector3(p.Pos.X, p.Pos.Y, p.Pos.Z+p.Vel.Z), rl.NewVector3(1, 2, 1), platform.Pos, platform.Size) {
			p.Vel.Z = 0
		}
	}

	p.Vel.Y = p.Fake_Y_Vel * rl.GetFrameTime()

	p.Pos = rl.Vector3Add(p.Pos, p.Vel)
}

var Player = NewPlayer(rl.NewVector3(0, 2, 0))
