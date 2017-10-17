package TurboOcto

import (
    "github.com/ByteArena/box2d"
	"github.com/veandco/go-sdl2/sdl"
    "github.com/Pixdigit/TurboOcto/internal/physicsEngine"
    "github.com/Pixdigit/TurboOcto/internal/renderEngine"
    "github.com/Pixdigit/TurboOcto/internal/eventEngine"
    "github.com/Pixdigit/TurboOcto/internal/sharedStructs"
    )

type Configuration struct {
    XGravitation, YGravitation float64
    Fullscreen bool
}

func CreateDefaultConfiguration() (Configuration) {
    return Configuration{0, -10, true}
}

type Environment struct {
    gi renderEngine.GraphicsInterface
    world box2d.B2World
    eventHandler eventEngine.EventEnv
}

func Update(env *Environment) (bool, []sdl.Event) {
    return env.eventHandler.UpdateEvents()
}

func LoadSphericalObject(filename string, env Environment) (sharedStructs.Sprite) {
    image := env.gi.LoadImage(filename, env.world)
	_, _, width, height, _ := image.Query()
	rect := &sdl.Rect{0, 0, width, height}

	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Position.Set(0, -10)
    bodyDef.Type = box2d.B2BodyType.B2_dynamicBody
	body := env.world.CreateBody(&bodyDef)
    shape := box2d.MakeB2CircleShape()
	shape.B2Shape.M_radius = 0.5
	body.CreateFixture(shape, 0.0)

    return sharedStructs.Sprite{Texture: image, Rect: rect, Body: body}
}



func CreateEnvironment(conf Configuration) (Environment) {
    gi := renderEngine.CreateGraphicsInterface(true)
    world := physicsEngine.CreateWorld(conf.XGravitation, conf.YGravitation)
    eventHandler := eventEngine.EventEnv{}
    return Environment{gi, world, eventHandler}
}

func (env Environment) Destroy() {
	env.gi.Renderer.Destroy()
	env.gi.Window.Destroy()
	sdl.Quit()
}

func Blit(sprite sharedStructs.Sprite, env Environment) {
    env.gi.Blit(sprite)
}

func Flip(env Environment) {
    env.gi.Flip()
}
