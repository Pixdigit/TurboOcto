package renderEngine

import "github.com/veandco/go-sdl2/sdl"

type RenderEngine struct {
	window        *sdl.Window
	Renderer      *sdl.Renderer
	ScreenTexture *sdl.Texture
}

func (r *RenderEngine) Init() {

}

func (r *RenderEngine) Destroy() {
	r.ScreenTexture.Destroy()
	r.Renderer.Destroy()
	r.window.Destroy()
}


func CreateRenderEngine(fullscreen bool) RenderEngine{
	renderEngine := RenderEngine{}
	const windowTitle string = "GoGame"
	var windowFlags uint32 = uint32(sdl.WINDOW_SHOWN)
	if fullscreen {
		windowFlags = windowFlags | uint32(sdl.WINDOW_FULLSCREEN_DESKTOP)
	}

	window, _ := sdl.CreateWindow(windowTitle, 0, 0, 0, 0, windowFlags)
	renderer, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_PRESENTVSYNC)
	screenTexture, _ := renderer.CreateTexture(uint32(sdl.PIXELFORMAT_ABGR8888), sdl.TEXTUREACCESS_STREAMING, 0, 0)

	renderEngine = RenderEngine{window: window, Renderer: renderer, ScreenTexture: screenTexture}

	return renderEngine
}

func main() {
}
