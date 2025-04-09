package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func collide(pos1, size1, pos2, size2 rl.Vector3) bool {
	return pos1.X+(size1.X) > pos2.X-(size2.X) &&
		pos1.X-(size1.X) < pos2.X+(size2.X) &&
		pos1.Y+(size1.Y) > pos2.Y-(size2.Y) &&
		pos1.Y-(size1.Y) < pos2.Y+(size2.Y) &&
		pos1.Z+(size1.Z) > pos2.Z-(size2.Z) &&
		pos1.Z-(size1.Z) < pos2.Z+(size2.Z)
}

func deg2rad(num float64) float64 {
	return num * 3.14159 / 180
}
