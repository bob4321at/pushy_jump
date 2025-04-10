package scenes

import (
	"encoding/json"
	"fmt"
	"image/color"
	"math"
	"os"
	"pushy/camera"
	platform "pushy/level"
	"pushy/lighting"
	"pushy/player"
	"pushy/utils"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
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
		if level_edit_mode {
			if selected_platform == i {
				rl.DrawCube(p.Pos, p.Size.X*2, p.Size.Y*2, p.Size.Z*2, rl.White)
			} else {
				p.Draw()
			}
		} else {
			p.Draw()
		}
	}

	rl.DrawModel(dog_model, rl.NewVector3(30, -2, 30), 1, rl.Beige)

	rl.EndMode3D()

	rl.DrawText("FPS: "+strconv.Itoa(int(rl.GetFPS())), 10, 10, 32, rl.Black)

	rl.DrawCircle(int32(rl.GetScreenWidth()/2), int32(rl.GetScreenHeight()/2), 5, rl.Black)

	rl.EndDrawing()
}

var lights = []lighting.Light{}

var level_edit_mode = false
var selected_platform = 0

func Game_Scene_Update() {
	if !level_edit_mode {
		player.Player.Update()
	} else {
		camera.Camera.FreeCam()
		if rl.IsKeyPressed(rl.KeyE) {
			if selected_platform+1 < len(platform.Platforms) {
				selected_platform += 1
			}
		}
		if rl.IsKeyPressed(rl.KeyQ) {
			if selected_platform-1 >= 0 {
				selected_platform -= 1
			}
		}
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) || rl.IsMouseButtonPressed(rl.MouseButtonRight) {
			_, hit, _, plat := utils.RaycastForPlatform(camera.Camera.Pos, camera.Camera.Rot, 1000)
			if hit {
				for pi := 0; pi < len(platform.Platforms); pi++ {
					if plat == &platform.Platforms[pi] {
						selected_platform = pi
					}
				}
			}
		}

		plat := &platform.Platforms[0]
		if selected_platform < len(platform.Platforms) {
			plat = &platform.Platforms[selected_platform]
		}

		if rl.IsKeyPressed(rl.KeyK) {
			plat.Pos.X += 1 * rl.GetFrameTime()
		} else if rl.IsKeyPressed(rl.KeyJ) {
			plat.Pos.X -= 1 * rl.GetFrameTime()
		}
		if rl.IsKeyPressed(rl.KeyL) {
			plat.Pos.Z += 1
		} else if rl.IsKeyPressed(rl.KeyH) {
			plat.Pos.Z -= 1
		}
		if rl.IsKeyPressed(rl.KeyK) {
			plat.Pos.X += 1
		} else if rl.IsKeyPressed(rl.KeyJ) {
			plat.Pos.X -= 1
		}
		if rl.IsKeyPressed(rl.KeyU) {
			plat.Pos.Y += 1
		} else if rl.IsKeyPressed(rl.KeyP) {
			plat.Pos.Y -= 1
		}

		if rl.IsKeyPressed(rl.KeyN) {
			plat.Size.X += 1
		} else if rl.IsKeyPressed(rl.KeyPeriod) {
			plat.Size.X -= 1
		}
		if rl.IsKeyPressed(rl.KeyM) {
			plat.Size.Z += 1
		} else if rl.IsKeyPressed(rl.KeyComma) {
			plat.Size.Z -= 1
		}
		if rl.IsKeyPressed(rl.KeyI) {
			plat.Size.Y += 1
		} else if rl.IsKeyPressed(rl.KeyO) {
			plat.Size.Y -= 1
		}

		if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
			utils.RemoveArrayElement(selected_platform, &platform.Platforms)
			if selected_platform > len(platform.Platforms) {
				selected_platform = len(platform.Platforms) - 1
			}
		}

		if rl.IsMouseButtonPressed(rl.MouseButtonMiddle) {
			hit_pos, hit, _ := utils.Raycast(camera.Camera.Pos, camera.Camera.Rot, 1000)
			if hit {
				platform.Platforms = append(platform.Platforms, platform.NewPlatform(hit_pos, rl.NewVector3(1, 1, 1)))
			} else {
				move_dir := rl.NewVector3(0, 0, 0)
				move_dir.X = float32(math.Cos(utils.Deg2Rad(float64(camera.Camera.Rot.Y)))) * 9
				move_dir.Y = float32(math.Sin(utils.Deg2Rad(float64(camera.Camera.Rot.X)))) * 9
				move_dir.Z = float32(math.Sin(utils.Deg2Rad(float64(camera.Camera.Rot.Y)))) * 9
				platform.Platforms = append(platform.Platforms, platform.NewPlatform(rl.Vector3Add(camera.Camera.Pos, move_dir), rl.NewVector3(1, 1, 1)))
			}
			_, hit, _, plat := utils.RaycastForPlatform(camera.Camera.Pos, camera.Camera.Rot, 1000)
			if hit {
				for pi := 0; pi < len(platform.Platforms); pi++ {
					if plat == &platform.Platforms[pi] {
						selected_platform = pi
					}
				}
			}
		}
	}

	if rl.IsKeyPressed(rl.KeyTab) {
		level_edit_mode = !level_edit_mode
		player.Player.Pos = camera.Camera.Pos
	}

	camera.Camera.Update()

	if rl.IsKeyPressed(rl.KeyV) {
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
