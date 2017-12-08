package TurboOcto

import (
    "github.com/ByteArena/box2d"
    "github.com/veandco/go-sdl2/sdl"
    "github.com/Pixdigit/TurboOcto/internal/physicsEngine"
    "github.com/Pixdigit/TurboOcto/internal/renderEngine"
    "github.com/Pixdigit/TurboOcto/internal/eventEngine"
    "github.com/pkg/errors"
)

type Environment struct {
    renderEngine *renderEngine.RenderEngine
    world *box2d.B2World
    eventHandler *eventEngine.EventEnv
    timestep float64
    positionAccuracy, velocityAccuracy int
}

func CreateEnvironment(conf Configuration) (Environment, error) {
    gi, err := renderEngine.NewRenderEngine(conf.ScreenWidth, conf.ScreenHeight, conf.Fullscreen)
    world := physicsEngine.CreateWorld(conf.XGravitation, conf.YGravitation)
    eventHandler := eventEngine.EventEnv{}
    if err != nil {return Environment{}, errors.Wrap(err, "Error while creating Environment")}
    return Environment{&gi, &world, &eventHandler, conf.Timestep, conf.PositionAccuracy, conf.VelocityAccuracy}, nil
}

func (env Environment) Destroy() (error){
	err := env.renderEngine.Renderer.Destroy()
    if err != nil {return errors.Wrap(err, "Could not destroy renderer")}
	err = env.renderEngine.Window.Destroy()
    if err != nil {return errors.Wrap(err, "Could not destroy window")}
	sdl.Quit()
    return nil
}

func (env Environment) Update() ([]sdl.Event) {
    env.world.Step(1.0 / 60, 6, 2)
    return env.eventHandler.UpdateEvents()
}
