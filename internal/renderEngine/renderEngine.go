package renderEngine

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ByteArena/box2d"
	"github.com/Pixdigit/TurboOcto/internal/sharedStructs"
	"github.com/pkg/errors"
	)

func Init() {
	sdl.InitSubSystem(0)
}

type RenderEngine struct {
	Renderer      *sdl.Renderer
	Window        *sdl.Window
}

func NewRenderEngine(windowWidth, windowHeight int32, fullscreen bool) (RenderEngine, error){
	const windowTitle string = "GoGame"
	var windowFlags uint32 = uint32(sdl.WINDOW_SHOWN)
	var window *sdl.Window
	var err error

	if (windowWidth != 0 && windowHeight != 0 && !fullscreen) {
		window, err = sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, windowWidth, windowHeight, windowFlags)
	} else if fullscreen {
		windowFlags = windowFlags | uint32(sdl.WINDOW_FULLSCREEN_DESKTOP)
		window, err = sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 0, 0, windowFlags)
	} else {
		windowFlags = windowFlags | uint32(sdl.WINDOW_MAXIMIZED) | uint32(sdl.WINDOW_RESIZABLE)
		window, err = sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, windowWidth, windowHeight, windowFlags)
		window.Maximize()
		sdl.Delay(500)
		width, height := window.GetMaximumSize()
		window.SetSize(width, height)
	}
	if err != nil {return RenderEngine{}, errors.Wrap(err, "Error creating RenderEngine")}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_PRESENTVSYNC)
	if err != nil {return RenderEngine{}, errors.Wrap(err, "Error creating RenderEngine")}

	rE := RenderEngine{Window: window, Renderer: renderer}
	return rE, nil
}

func (r *RenderEngine) LoadImage(path string, world *box2d.B2World) (*sdl.Texture, error){
	i, err := sdl.LoadBMP(path)
	if err != nil {return &sdl.Texture{}, errors.Wrap(err, "Could not load imagefile " + path)}
    image, err := r.Renderer.CreateTextureFromSurface(i)
	if err != nil {return &sdl.Texture{}, errors.Wrap(err, "Could not create texture from image")}
	return image, nil
}

func (r *RenderEngine) Blit(sprite sharedStructs.Sprite) (error) {
	err := r.Renderer.Copy(sprite.Texture, nil, sprite.Rect)
	if err != nil {return errors.Wrap(err, "Error while Blitting " + sprite.Name)}
	return nil
}

func (r *RenderEngine) Flip() (error) {
	r.Renderer.Present()
	err := r.Renderer.Clear()
	if err != nil {return errors.Wrap(err, "Error while flipping")}
	return nil
}
