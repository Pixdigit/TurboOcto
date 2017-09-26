package renderEngine

import "github.com/veandco/go-sdl2/sdl"

type RenderEngine struct {
	window        *sdl.Window
	Renderer      *sdl.Renderer
	ScreenTexture *sdl.Texture
}

func (r *RenderEngine) Init() {

	var windowRect sdl.Rect = sdl.Rect{sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 0, 0}
	const window_title string = "Test"
	var window_flags uint32 = uint32(sdl.WINDOW_SHOWN) | uint32(sdl.WINDOW_RESIZABLE) | uint32(sdl.WINDOW_FULLSCREEN_DESKTOP)

	sdl.Init(sdl.INIT_EVERYTHING)
	//sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "linear")

	window, _ := sdl.CreateWindow(window_title, int(windowRect.X), int(windowRect.Y), int(windowRect.W), int(windowRect.H), window_flags)
	defer window.Destroy()
	w, h := window.GetSize()
	windowRect.W = int32(w)
	windowRect.H = int32(h)
	windowRect.X = 0
	windowRect.Y = 0
	renderer, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_PRESENTVSYNC) // -1: use first availible driver 0: No flags
	defer renderer.Destroy()
	screenTexture, _ := renderer.CreateTexture(uint32(sdl.PIXELFORMAT_ABGR8888), sdl.TEXTUREACCESS_STREAMING, int(windowRect.W), int(windowRect.H))

	_ = sdl.ShowCursor(sdl.DISABLE)

	defer screenTexture.Destroy()
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
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
