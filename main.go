package main

import (
	"pushy/scenes"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func update() {
	if !scenes.List_Of_Scenes[scenes.Current_Scene].Setup_run {
		scenes.List_Of_Scenes[scenes.Current_Scene].Setup_run = true
	}

	scenes.List_Of_Scenes[scenes.Current_Scene].Update()
}

func draw() {
	scenes.List_Of_Scenes[scenes.Current_Scene].Draw()
}

func main() {
	rl.InitWindow(1280, 720, "push jump")

	// rl.SetTargetFPS(60)

	rl.SetConfigFlags(rl.FlagMsaa4xHint)

	rl.DisableCursor()
	scenes.List_Of_Scenes[scenes.Current_Scene].Setup()

	for !rl.WindowShouldClose() {
		update()
		draw()
	}

	rl.CloseWindow()
}
