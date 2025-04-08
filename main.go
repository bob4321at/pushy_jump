package main

import (
	"image/color"

	"github.com/gen2brain/raylib-go/raylib"
)

var light_shader rl.Shader

var lights = []Light{}

var (
	cube_mesh rl.Model
)

func update() {
	camera.Update()

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

	rl.EndMode3D()

	rl.EndDrawing()
}

func main() {
	rl.InitWindow(1280, 720, "push jump")

	cube_mesh = rl.LoadModel("./cube.obj")

	light_shader = rl.LoadShader("./shaders/lighting.vs", "./shaders/lighting.fs")

	*light_shader.Locs = rl.GetShaderLocation(light_shader, "viewPos")

	ambientLoc := rl.GetShaderLocation(light_shader, "ambient")
	shaderValue := []float32{0.1, 0.1, 0.1, 1.0}
	rl.SetShaderValue(light_shader, ambientLoc, shaderValue, rl.ShaderUniformVec4)
	lights = []Light{
		NewLight(LightTypePoint, rl.NewVector3(0, 100, 100), rl.NewVector3(0, 0, 0), rl.White, light_shader),
	}

	rl.DisableCursor()

	cube_mesh.Materials.Shader = light_shader
	cube_mesh.Materials.Maps.Texture = rl.LoadTexture("./texture.png")

	for !rl.WindowShouldClose() {
		update()
		draw()
	}

	rl.UnloadShader(light_shader)

	rl.CloseWindow()
}
