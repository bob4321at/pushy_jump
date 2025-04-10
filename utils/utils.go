package utils

import (
	"math"
	platform "pushy/level"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Collide(pos1, size1, pos2, size2 rl.Vector3) bool {
	return pos1.X+(size1.X) > pos2.X-(size2.X) &&
		pos1.X-(size1.X) < pos2.X+(size2.X) &&
		pos1.Y+(size1.Y) > pos2.Y-(size2.Y) &&
		pos1.Y-(size1.Y) < pos2.Y+(size2.Y) &&
		pos1.Z+(size1.Z) > pos2.Z-(size2.Z) &&
		pos1.Z-(size1.Z) < pos2.Z+(size2.Z)
}

func Raycast(pos, rot rl.Vector3, distance float32) (rl.Vector3, bool, float32) {
	for l := float32(0); l < distance; l++ {
		direction := rl.NewVector3(0, 0, 0)

		direction.X = float32(math.Cos(Deg2Rad(float64(rot.Y)))) * float32(math.Cos(Deg2Rad(float64(rot.X))))
		direction.Y = float32(math.Sin(Deg2Rad(float64(rot.X))))
		direction.Z = float32(math.Sin(Deg2Rad(float64(rot.Y)))) * float32(math.Cos(Deg2Rad(float64(rot.X))))

		for pi := 0; pi < len(platform.Platforms); pi++ {
			platform := &platform.Platforms[pi]
			if Collide(rl.Vector3Add(pos, rl.NewVector3(direction.X*l, direction.Y*l, direction.Z*l)), rl.NewVector3(0.001, 0.001, 0.001), platform.Pos, platform.Size) {
				return rl.Vector3Add(pos, rl.NewVector3(direction.X*l, direction.Y*l, direction.Z*l)), true, l
			}
		}
	}

	return rl.NewVector3(0, 0, 0), false, -1
}

func Deg2Rad(num float64) float64 {
	return num * 3.14159 / 180
}
