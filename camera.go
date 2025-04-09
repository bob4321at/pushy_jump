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

var camera = NewCamera(rl.NewVector3(5, 0, 0), rl.NewVector3(0, -90, 0))
