package renderEngine

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ByteArena/box2d"
	"github.com/Pixdigit/TurboOcto/internal/sharedStructs"
	)

func Init(env box2d.B2World) {
	sdl.InitSubSystem(0)
}

type GraphicsInterface struct {
	Window        *sdl.Window
	Renderer      *sdl.Renderer
}

func CreateGraphicsInterface(fullscreen bool) (GraphicsInterface){
	gi := GraphicsInterface{}
	const windowTitle string = "GoGame"
	var windowFlags uint32 = uint32(sdl.WINDOW_SHOWN)
	if fullscreen {
		windowFlags = windowFlags | uint32(sdl.WINDOW_FULLSCREEN_DESKTOP)
	}

	window, _ := sdl.CreateWindow(windowTitle, 0, 0, 0, 0, windowFlags)
	renderer, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_PRESENTVSYNC)

	gi = GraphicsInterface{Window: window, Renderer: renderer}

	return gi
}

func (r *GraphicsInterface) LoadImage(path string, world box2d.B2World) (*sdl.Texture){
	i, _ := sdl.LoadBMP("./assets/images/test1.bmp")
    image, _ := r.Renderer.CreateTextureFromSurface(i)
	return image
}

func (r *GraphicsInterface) Blit(sprite sharedStructs.Sprite) {
	r.Renderer.Copy(sprite.Texture, nil, sprite.Rect)
}

func (r *GraphicsInterface) Flip() {
	r.Renderer.Present()
	r.Renderer.Clear()
}
