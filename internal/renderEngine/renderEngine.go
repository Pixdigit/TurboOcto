package renderEngine

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ByteArena/box2d"
	"github.com/Pixdigit/TurboOcto/internal/sharedStructs"
	"fmt"
	)

func Init() {
	sdl.InitSubSystem(0)
}

type GraphicsInterface struct {
	Window        *sdl.Window
	Renderer      *sdl.Renderer
}

func CreateGraphicsInterface(windowWidth, windowHeight int, fullscreen bool) (GraphicsInterface){
	gi := GraphicsInterface{}
	const windowTitle string = "GoGame"
	var windowFlags uint32 = uint32(sdl.WINDOW_SHOWN)
	var window *sdl.Window

	if (windowWidth != 0 && windowHeight != 0 && !fullscreen) {
		window, _ = sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, windowWidth, windowHeight, windowFlags)
	} else if fullscreen {
		windowFlags = windowFlags | uint32(sdl.WINDOW_FULLSCREEN_DESKTOP)
		window, _ = sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 0, 0, windowFlags)
	} else {
		fmt.Println("Test")
		windowFlags = windowFlags | uint32(sdl.WINDOW_MAXIMIZED) | uint32(sdl.WINDOW_RESIZABLE)
		windowWidth, windowHeight = 150, 150
		window, _ = sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, windowWidth, windowHeight, windowFlags)
		window.Maximize()
		//TODO: uncomment after implemented in go-sdl2
		//window.Resizable(false)

	}

	renderer, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_PRESENTVSYNC)

	gi = GraphicsInterface{Window: window, Renderer: renderer}

	return gi
}

func (r *GraphicsInterface) LoadImage(path string, world *box2d.B2World) (*sdl.Texture){
	i, _ := sdl.LoadBMP(path)
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
