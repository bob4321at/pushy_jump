package scenes

import (
	"encoding/json"
	"fmt"
	"image/color"
	"os"
	"pushy/camera"
	platform "pushy/level"
	"pushy/lighting"
	"pushy/player"

	"github.com/gen2brain/raylib-go/raylib"
)

var Game_Scene = NewScene(1, Game_Scene_Draw, Game_Scene_Update, Game_Scene_Setup)

var light_shader rl.Shader

func Game_Scene_Setup() {
	f, err := os.ReadFile("level.json")
	if err != nil {
		platform.Platforms = []platform.Platform{
			platform.NewPlatform(rl.NewVector3(0, -5, 0), rl.NewVector3(1000, 1, 1000)),
		}
	} else {
		temp_var := []platform.Platform{}
		err := json.Unmarshal(f, &temp_var)
		if err != nil {
			panic(err)
		}
		platform.Platforms = temp_var
	}

	platform.Cube_Mesh = rl.LoadModel("./cube.obj")
	dog_model = rl.LoadModel("./models/dog.obj")

	light_shader = rl.LoadShader("./shaders/lighting.vs", "./shaders/lighting.fs")

	*light_shader.Locs = rl.GetShaderLocation(light_shader, "viewPos")

	ambientLoc := rl.GetShaderLocation(light_shader, "ambient")
	shaderValue := []float32{0.1, 0.1, 0.1, 1.0}
	rl.SetShaderValue(light_shader, ambientLoc, shaderValue, rl.ShaderUniformVec4)
	lights = []lighting.Light{
		lighting.NewLight(rl.NewVector3(0, 100, 100), rl.NewVector3(0, 0, 0), rl.White, light_shader),
	}
	platform.Cube_Mesh.Materials.Shader = light_shader
	platform.Cube_Mesh.Materials.Maps.Texture = rl.LoadTexture("./texture.png")

	dog_model.Materials.Shader = light_shader
}

var (
	dog_model rl.Model
)

func Game_Scene_Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(color.RGBA{0, 125, 255, 255})

	rl.BeginMode3D(camera.Camera.Camera)

	for i := 0; i < len(platform.Platforms); i++ {
		p := platform.Platforms[i]
		p.Draw()
	}

	rl.DrawModel(dog_model, rl.NewVector3(30, -2, 30), 1, rl.Beige)

	rl.EndMode3D()

	rl.DrawCircle(int32(rl.GetScreenWidth()/2), int32(rl.GetScreenHeight()/2), 5, rl.Black)

	rl.EndDrawing()
}

var lights = []lighting.Light{}

var level_edit_mode = false

func Game_Scene_Update() {
	if !level_edit_mode {
		player.Player.Update()
	} else {
		camera.Camera.FreeCam()
	}

	if rl.IsKeyPressed(rl.KeyTab) {
		level_edit_mode = !level_edit_mode
		player.Player.Pos = camera.Camera.Pos
	}

	camera.Camera.Update()

	if rl.IsKeyPressed(rl.KeyE) {
		tesst, err := json.Marshal(platform.Platforms)
		if err != nil {
			panic(err)
		}

		f, err := os.Create("level.json")
		if err != nil {
			panic(err)
		}
		f.Write(tesst)

		test_var := []platform.Platform{}

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
