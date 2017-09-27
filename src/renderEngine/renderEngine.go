package renderEngine

import "github.com/veandco/go-sdl2/sdl"

type RenderEngine struct {
	window        *sdl.Window
	Renderer      *sdl.Renderer
}

func (r *RenderEngine) Init() {
}

func (r *RenderEngine) LoadImage(path string) (*sdl.Texture){
	i, _ := sdl.LoadBMP("./assets/images/test1.bmp")
    image, _ := r.Renderer.CreateTextureFromSurface(i)
	return image
}

func (r *RenderEngine) Blit(texture *sdl.Texture, dest *sdl.Rect) {
	r.Renderer.Copy(texture, nil, dest)
}

func (r *RenderEngine) Flip() {
	r.Renderer.Present()
	r.Renderer.Clear()
}

func (r *RenderEngine) Destroy() {
	r.Renderer.Destroy()
	r.window.Destroy()
	sdl.Quit()
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

	renderEngine = RenderEngine{window: window, Renderer: renderer}

	return renderEngine
}

func main() {
}
