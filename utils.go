package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func collide(pos1, size1, pos2, size2 rl.Vector3) bool {
	return pos1.X+(size1.X) > pos2.X-(size2.X) &&
		pos1.X-(size1.X) < pos2.X+(size2.X) &&
		pos1.Y+(size1.Y) > pos2.Y-(size2.Y) &&
		pos1.Y-(size1.Y) < pos2.Y+(size2.Y) &&
		pos1.Z+(size1.Z) > pos2.Z-(size2.Z) &&
		pos1.Z-(size1.Z) < pos2.Z+(size2.Z)
}

func raycast(pos, rot rl.Vector3, distance float32) (rl.Vector3, bool) {
	for l := float32(0); l < distance; l++ {
		direction := rl.NewVector3(0, 0, 0)

		direction.X = float32(math.Cos(deg2rad(float64(rot.Y)))) * float32(math.Cos(deg2rad(float64(rot.X))))
		direction.Y = float32(math.Sin(deg2rad(float64(rot.X))))
		direction.Z = float32(math.Sin(deg2rad(float64(rot.Y)))) * float32(math.Cos(deg2rad(float64(rot.X))))

		for pi := 0; pi < len(platforms); pi++ {
			platform := &platforms[pi]
			if collide(rl.Vector3Add(pos, rl.NewVector3(direction.X*l, direction.Y*l, direction.Z*l)), rl.NewVector3(0.001, 0.001, 0.001), platform.Pos, platform.Size) {
				return rl.Vector3Add(pos, rl.NewVector3(direction.X*l, direction.Y*l, direction.Z*l)), true
			}
		}
	}

	return rl.NewVector3(0, 0, 0), false
}

func deg2rad(num float64) float64 {
	return num * 3.14159 / 180
}
