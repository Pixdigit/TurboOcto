package turboOcto

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"gitlab.com/Pixdigit/geometry"
)

var screenSize *geometry.Size
var windowSize *geometry.Size
var canvasSize *geometry.Size
var vSize *geometry.Size
var frameCount int32
var isFullscreen bool

var screenRenderer *sdl.Renderer
var window *sdl.Window
var displayIndex int //TODO: Dynamically update when window moved

var rmask uint32 = 0x000000ff
var gmask uint32 = 0x0000ff00
var bmask uint32 = 0x00ff0000
var amask uint32 = 0xff000000

func initializeGraphics() (err error) {
	//Default is LIL_ENDIAN
	if sdl.BYTEORDER == sdl.BIG_ENDIAN {
		rmask = 0xff000000
		gmask = 0x00ff0000
		bmask = 0x0000ff00
		amask = 0x000000ff
	}

	//Create graphical interfaces
	windowFlags := uint32(sdl.WINDOW_SHOWN) | uint32(sdl.WINDOW_FULLSCREEN_DESKTOP)
	window, err = sdl.CreateWindow("", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 0, 0, windowFlags);	if err != nil {return errors.Wrap(err, "could not create window")}
	window.SetGrab(true)

	screenRenderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_PRESENTVSYNC);	if err != nil {return errors.Wrap(err, "could not create screenRenderer")}

	displayIndex, err := window.GetDisplayIndex();	if err != nil {return errors.Wrap(err, "could not get display index")}
	dmode, err := sdl.GetDesktopDisplayMode(displayIndex);	if err != nil {return errors.Wrap(err, "could not get display mode")}

	//initialize sizes
	screenSize = &geometry.Size{geometry.Scalar(dmode.W), geometry.Scalar(dmode.H)}
	w, h := window.GetSize()
	windowSize = &geometry.Size{geometry.Scalar(w), geometry.Scalar(h)}
	w, h, err = screenRenderer.GetOutputSize();	if err != nil {return errors.Wrap(err, "could not read output size")}
	canvasSize = &geometry.Size{geometry.Scalar(w), geometry.Scalar(h)}
	vSize = canvasSize.Copy()

	if isFullscreen {
		Fullscreen()
	} else {
		Windowed()
		SetWindowSize(*screenSize.GetScaled(1 / 4.0))
	}

	Clear()
	return nil
}

func getSDLSize(size *geometry.Size) (int32, int32) {
	return int32(size.Width), int32(size.Height)
}

func SetDecoration(title string, iconPath string) error {
	window.SetTitle(title)
	iconPath = Cfg.ResourcePath + iconPath
	if iconPath != "" {
		if exists, err := pathExists(iconPath); err != nil {
			return errors.Wrap(err, "could not check wether icon file exists")
		} else if !exists {
			return errors.New("path to icon does not exist")
		} else {
			icon, err := img.Load(iconPath);	if err != nil {return errors.Wrap(err, "could not load icon from path")}
			window.SetIcon(icon)
		}
	}
	return nil
}

func Fullscreen() error {
	window.SetSize(getSDLSize(screenSize))
	window.SetFullscreen(sdl.WINDOW_FULLSCREEN)
	canvasSize = screenSize.Copy()
	FillScreen(0, 0, 0, 0)
	Clear()
	isFullscreen = true
	return nil
}
func Windowed() error {
	const SDL_WINDOW_WINDOWED = 0
	window.SetFullscreen(SDL_WINDOW_WINDOWED)
	window.SetSize(getSDLSize(windowSize))
	window.SetPosition(sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED)
	canvasSize = windowSize.Copy()
	Clear()
	isFullscreen = false
	return nil
}

func SetVirtualSize(size geometry.Size) error {
	vSize = &size
	screenRenderer.SetLogicalSize(getSDLSize(&size))
	err := Clear();	if err != nil {return errors.Wrap(err, "could not clear window after changing virtual size")}
	return nil
}
func SetWindowSize(size geometry.Size) error {
	windowSize = size.Copy()
	if !isFullscreen {
		window.SetSize(getSDLSize(windowSize))
	}
	return nil
}
func WindowSize() (geometry.Size, error) {
	return *windowSize.Copy(), nil
}
func VSize() (geometry.Size, error) {
	return *vSize.Copy(), nil
}

func FillScreen(r, g, b, a uint8) error {
	err := screenRenderer.SetDrawColor(r, g, b, a);	if err != nil {return errors.Wrap(err, "could not set draw color for fill operation")}
	err = screenRenderer.FillRect(nil);	if err != nil {return errors.Wrap(err, "could not execute fill operation")}
	return nil
}
func Clear() error {
	return FillScreen(0, 0, 0, 0)
}
func Render() error {
	for _, elem := range zSpace.Elems() {
		//TODO: check for errors
		elem.(RenderElement).Render()
		switch thing := elem.(type) {
		case *Sprite:
			thing.update()
		}
	}

	screenRenderer.Present()
	frameCount += 1
	//Clear up dirty frameBuffer
	err := Clear();	if err != nil {return errors.Wrap(err, "could not prepare next frame")}
	return nil
}
