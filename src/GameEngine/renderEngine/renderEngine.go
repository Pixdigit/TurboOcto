package renderEngine

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ByteArena/box2d"
	)

var Env box2d.B2World

func Init(env box2d.B2World) {
	Env = env
	sdl.InitSubSystem(0)
}

type Sprite struct {
	Texture *sdl.Texture
	Rect *sdl.Rect
	Body *box2d.B2Body
}

type GraphicsInterface struct {
	Window        *sdl.Window
	Renderer      *sdl.Renderer
}

func (r *GraphicsInterface) LoadImage(path string, world box2d.B2World) (Sprite){
	i, _ := sdl.LoadBMP("./assets/images/test1.bmp")
    image, _ := r.Renderer.CreateTextureFromSurface(i)

	_, _, width, height, _ := image.Query()
	rect := &sdl.Rect{0, 0, width, height}

	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Position.Set(0, -10)
	body := world.CreateBody(&bodyDef)
	shape := box2d.MakeB2CircleShape()
	shape.B2Shape.M_radius = 0.5
	body.CreateFixture(shape, 0.0)

	return Sprite{image, rect, body}
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

func (r *GraphicsInterface) Blit(sprite Sprite) {
	r.Renderer.Copy(sprite.Texture, nil, sprite.Rect)
}

func (r *GraphicsInterface) Flip() {
	r.Renderer.Present()
	r.Renderer.Clear()
}
