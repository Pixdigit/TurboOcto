package TurboOcto

import (
    "github.com/ByteArena/box2d"
    "github.com/veandco/go-sdl2/sdl"
    "github.com/Pixdigit/TurboOcto/internal/physicsEngine"
    "github.com/Pixdigit/TurboOcto/internal/renderEngine"
    "github.com/Pixdigit/TurboOcto/internal/eventEngine"
)

type Environment struct {
    renderEngine *renderEngine.RenderEngine
    world *box2d.B2World
    eventHandler *eventEngine.EventEnv
    timestep float64
    positionAccuracy, velocityAccuracy int
}

func CreateEnvironment(conf Configuration) (Environment) {
    gi := renderEngine.NewRenderEngine(conf.ScreenWidth, conf.ScreenHeight, conf.Fullscreen)
    world := physicsEngine.CreateWorld(conf.XGravitation, conf.YGravitation)
    eventHandler := eventEngine.EventEnv{}
    return Environment{&gi, &world, &eventHandler, conf.Timestep, conf.PositionAccuracy, conf.VelocityAccuracy}
}

func (env Environment) Destroy() {
	env.renderEngine.Renderer.Destroy()
	env.renderEngine.Window.Destroy()
	sdl.Quit()
}

func (env Environment) Update() ([]sdl.Event) {
    env.world.Step(1.0 / 60, 6, 2)
    return env.eventHandler.UpdateEvents()
}
