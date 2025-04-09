package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var light_shader rl.Shader

var lights = []Light{}

var (
	cube_mesh rl.Model
	dog_model rl.Model

	dt float64
)

var level_edit_mode = false

func update() {
	if !level_edit_mode {
		player.Update()
	} else {
		camera.FreeCam()
	}

	if rl.IsKeyPressed(rl.KeyTab) {
		level_edit_mode = !level_edit_mode
		player.Pos = camera.Pos
	}

	camera.Update()

	if rl.IsKeyPressed(rl.KeyE) {
		tesst, err := json.Marshal(platforms)
		if err != nil {
			panic(err)
		}

		f, err := os.Create("level.json")
		if err != nil {
			panic(err)
		}
		f.Write(tesst)

		test_var := []Platform{}

		err = json.Unmarshal(tesst, &test_var)
		if err != nil {
			panic(err)
		}

		fmt.Println(test_var)
	}

	for li := 0; li < len(lights); li++ {
		light := lights[li]
		light.UpdateValues()
	}
}

func draw() {
	rl.BeginDrawing()

	rl.ClearBackground(color.RGBA{0, 125, 255, 255})

	rl.BeginMode3D(camera.Camera)

	for i := 0; i < len(platforms); i++ {
		p := platforms[i]
		p.Draw()
	}

	rl.DrawModel(dog_model, rl.NewVector3(30, -2, 30), 1, rl.Beige)

	rl.EndMode3D()

	rl.DrawCircle(int32(rl.GetScreenWidth()/2), int32(rl.GetScreenHeight()/2), 5, rl.Black)

	rl.EndDrawing()
}

func main() {
	rl.InitWindow(1280, 720, "push jump")

	// rl.SetTargetFPS(60)

	rl.SetConfigFlags(rl.FlagMsaa4xHint)

	f, err := os.ReadFile("level.json")
	if err != nil {
		platforms = []Platform{
			NewPlatform(rl.NewVector3(0, -5, 0), rl.NewVector3(100, 1, 100)),
		}
	} else {
		temp_var := []Platform{}
		err := json.Unmarshal(f, &temp_var)
		if err != nil {
			panic(err)
		}
		platforms = temp_var
	}

	cube_mesh = rl.LoadModel("./cube.obj")
	dog_model = rl.LoadModel("./models/dog.obj")

	light_shader = rl.LoadShader("./shaders/lighting.vs", "./shaders/lighting.fs")

	*light_shader.Locs = rl.GetShaderLocation(light_shader, "viewPos")

	ambientLoc := rl.GetShaderLocation(light_shader, "ambient")
	shaderValue := []float32{0.1, 0.1, 0.1, 1.0}
	rl.SetShaderValue(light_shader, ambientLoc, shaderValue, rl.ShaderUniformVec4)
	lights = []Light{
		NewLight(rl.NewVector3(0, 100, 100), rl.NewVector3(0, 0, 0), rl.White, light_shader),
	}

	rl.DisableCursor()

	cube_mesh.Materials.Shader = light_shader
	cube_mesh.Materials.Maps.Texture = rl.LoadTexture("./texture.png")

	dog_model.Materials.Shader = light_shader

	for !rl.WindowShouldClose() {
		update()
		draw()
	}

	rl.UnloadShader(light_shader)

	rl.CloseWindow()
}
