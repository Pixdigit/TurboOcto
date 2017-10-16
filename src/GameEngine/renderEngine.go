package GameEngine

import ("github.com/veandco/go-sdl2/sdl"
	"github.com/ByteArena/box2d")

type graphicsInterface struct {
	window        *sdl.Window
	renderer      *sdl.Renderer
}

func (r *graphicsInterface) LoadImage(path string) (Sprite){
	i, _ := sdl.LoadBMP("./assets/images/test1.bmp")
    image, _ := r.renderer.CreateTextureFromSurface(i)
	_, _, width, height, _ := image.Query()
	rect := &sdl.Rect{0, 0, width, height}
	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Position.Set(0, -10)
	body := World.CreateBody(&bodyDef)
	shape := box2d.MakeB2CircleShape()
	shape.B2Shape.M_radius = 0.5
	body.CreateFixture(shape, 0.0)

	return Sprite{image, rect, body}
}

func (r *graphicsInterface) Blit(sprite Sprite) {
	r.renderer.Copy(sprite.Texture, nil, sprite.Rect)
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

type Sprite struct {
	Texture *sdl.Texture
	Rect *sdl.Rect
	Body *box2d.B2Body
}
