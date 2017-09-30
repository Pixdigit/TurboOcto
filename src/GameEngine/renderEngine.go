package GameEngine

import "github.com/veandco/go-sdl2/sdl"

type graphicsInterface struct {
	window        *sdl.Window
	renderer      *sdl.Renderer
}

func (r *graphicsInterface) LoadImage(path string) (*sdl.Texture){
	i, _ := sdl.LoadBMP("./assets/images/test1.bmp")
    image, _ := r.renderer.CreateTextureFromSurface(i)
	return image
}

func (r *graphicsInterface) Blit(texture *sdl.Texture, dest *sdl.Rect) {
	r.renderer.Copy(texture, nil, dest)
}

func (r *graphicsInterface) Flip() {
	r.renderer.Present()
	r.renderer.Clear()
}

func (r *graphicsInterface) Destroy() {
	r.renderer.Destroy()
	r.window.Destroy()
	sdl.Quit()
}

func CreateGraphicsInterface(fullscreen bool) (graphicsInterface){
	gi := graphicsInterface{}
	const windowTitle string = "GoGame"
	var windowFlags uint32 = uint32(sdl.WINDOW_SHOWN)
	if fullscreen {
		windowFlags = windowFlags | uint32(sdl.WINDOW_FULLSCREEN_DESKTOP)
	}

	window, _ := sdl.CreateWindow(windowTitle, 0, 0, 0, 0, windowFlags)
	renderer, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_PRESENTVSYNC)

	gi = graphicsInterface{window: window, renderer: renderer}

	return gi
}
