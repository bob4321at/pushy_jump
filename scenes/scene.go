package scenes

type Scene struct {
	Id        int
	Draw      func()
	Update    func()
	Setup     func()
	Setup_run bool
}

func NewScene(id int, draw func(), update func(), setup func()) (scene Scene) {
	scene.Id = id

	scene.Draw = draw
	scene.Update = update
	scene.Setup = setup

	scene.Setup_run = false

	return scene
}

var List_Of_Scenes = []Scene{Game_Scene}
var Current_Scene = 0
var Old_Scene = Current_Scene
