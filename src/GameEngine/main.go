package GameEngine

import (
    "github.com/ByteArena/box2d"
	"github.com/veandco/go-sdl2/sdl"
    "./physicsEngine"
    "./renderEngine"
    )


func Update() {
    UpdateEvents()
}

func LoadSphericalObject(filename string, env Environment) (renderEngine.Sprite) {
    sprite := env.gi.LoadImage(filename)
    return sprite
}

type Configuration struct {
    XGravitation, YGravitation float64
    Fullscreen bool
}

func CreateDefaultCOnfiguration() (Configuration) {
    return Configuration{0, -10, true}
}

type Environment struct {
    gi renderEngine.GraphicsInterface
    world box2d.B2World
}

func CreateEnvironment(conf Configuration) (Environment) {
    gi := renderEngine.CreateGraphicsInterface(true)
    world := physicsEngine.CreateWorld(conf.XGravitation, conf.YGravitation)
    return Environment{gi, world}
}

func (env Environment) Destroy() {
	env.gi.Renderer.Destroy()
	env.gi.Window.Destroy()
	sdl.Quit()
}

func Blit(sprite renderEngine.Sprite, env Environment) {
    env.gi.Blit(sprite)
}

func Flip(env Environment) {
    env.gi.Flip()
}
