package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"
)

var OldMousePos = rl.NewVector2(0, 0)

type Camera struct {
	Pos         rl.Vector3
	Rot         rl.Vector3
	Sensitivity float32
	Camera      rl.Camera
}

func NewCamera(pos, rot rl.Vector3) (camera Camera) {
	camera.Pos = pos
	camera.Rot = rot
	camera.Camera = rl.NewCamera3D(pos, rl.NewVector3(0, 0, 0), rl.NewVector3(0, 1, 0), 66, rl.CameraPerspective)

	camera.Sensitivity = 0.3

	return camera
}

func (camera *Camera) Update() {
	direction := rl.NewVector3(0, 0, 0)

	direction.X = float32(math.Cos(deg2rad(float64(camera.Rot.Y)))) * float32(math.Cos(deg2rad(float64(camera.Rot.X))))
	direction.Y = float32(math.Sin(deg2rad(float64(camera.Rot.X))))
	direction.Z = float32(math.Sin(deg2rad(float64(camera.Rot.Y)))) * float32(math.Cos(deg2rad(float64(camera.Rot.X))))

	rl.Vector3Normalize(direction)

	camera.Camera.Position = camera.Pos
	camera.Camera.Target = rl.Vector3Add(direction, camera.Pos)

	camera.Rot.Y += (float32(rl.GetMouseX()) - OldMousePos.X) * camera.Sensitivity
	camera.Rot.X -= (float32(rl.GetMouseY()) - OldMousePos.Y) * camera.Sensitivity

	if rl.IsKeyDown(rl.KeyL) {
		camera.Rot.Y += 0.1
	} else if rl.IsKeyDown(rl.KeyH) {
		camera.Rot.Y -= 0.1
	}

	if rl.IsKeyDown(rl.KeyK) {
		camera.Rot.X += 0.1
	} else if rl.IsKeyDown(rl.KeyJ) {
		camera.Rot.X -= 0.1
	}

	if camera.Rot.X >= 90 {
		camera.Rot.X = 90
	} else if camera.Rot.X <= -90 {
		camera.Rot.X = -90
	}

	OldMousePos = rl.GetMousePosition()
}

func (camera *Camera) FreeCam() {
	move_dir := rl.NewVector3(0, 0, 0)
	move_dir.X = float32(math.Cos(deg2rad(float64(camera.Rot.Y)))) * rl.GetFrameTime() * 60
	move_dir.Y = float32(math.Sin(deg2rad(float64(camera.Rot.X)))) * rl.GetFrameTime() * 60
	move_dir.Z = float32(math.Sin(deg2rad(float64(camera.Rot.Y)))) * rl.GetFrameTime() * 60
	move_dir_side := rl.NewVector3(0, 0, 0)
	move_dir_side.X = float32(math.Cos(deg2rad(float64(camera.Rot.Y+90)))) * rl.GetFrameTime() * 60
	move_dir_side.Z = float32(math.Sin(deg2rad(float64(camera.Rot.Y+90)))) * rl.GetFrameTime() * 60

	if rl.IsKeyDown(rl.KeyW) {
		camera.Pos = rl.Vector3Add(rl.NewVector3(move_dir.X, move_dir.Y, move_dir.Z), camera.Pos)
	} else if rl.IsKeyDown(rl.KeyS) {
		camera.Pos = rl.Vector3Add(rl.NewVector3(-move_dir.X, -move_dir.Y, -move_dir.Z), camera.Pos)
	}

	if rl.IsKeyDown(rl.KeyD) {
		camera.Pos = rl.Vector3Add(rl.NewVector3(move_dir_side.X, move_dir_side.Y, move_dir_side.Z), camera.Pos)
	} else if rl.IsKeyDown(rl.KeyA) {
		camera.Pos = rl.Vector3Add(rl.NewVector3(-move_dir_side.X, -move_dir_side.Y, -move_dir_side.Z), camera.Pos)
	}
}

var camera = NewCamera(rl.NewVector3(5, 0, 0), rl.NewVector3(0, -90, 0))
