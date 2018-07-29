package turboOcto

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

var screenWidth, screenHeight int32
var windowWidth, windowHeight int32
var canvasWidth, canvasHeight int32
var vWidth, vHeight int32
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

	screenRenderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_PRESENTVSYNC);	if err != nil {return errors.Wrap(err, "could not create screenRenderer")}

	displayIndex, err := window.GetDisplayIndex();	if err != nil {return errors.Wrap(err, "could not get display index")}
	dmode, err := sdl.GetDesktopDisplayMode(displayIndex);	if err != nil {return errors.Wrap(err, "could not get display mode")}

	//initialize sizes
	screenWidth, screenHeight = dmode.W, dmode.H
	windowWidth, windowHeight = window.GetSize()
	canvasWidth, canvasHeight, err = screenRenderer.GetOutputSize();	if err != nil {return errors.Wrap(err, "could not read output size")}
	vWidth, vHeight = canvasWidth, canvasHeight

	if isFullscreen {
		Fullscreen()
	} else {
		Windowed()
		SetWindowSize(screenWidth/4, screenHeight/4)
	}

	Clear()
	return nil
}

func SetDecoration(title string, iconPath string) error {
	window.SetTitle(title)
	iconPath = "./assets/sprites/" + iconPath
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
	window.SetSize(screenWidth, screenHeight)
	window.SetFullscreen(sdl.WINDOW_FULLSCREEN)
	canvasWidth, canvasHeight = screenWidth, screenHeight
	FillScreen(0, 0, 0, 0)
	Clear()
	isFullscreen = true
	return nil
}
func Windowed() error {
	const SDL_WINDOW_WINDOWED = 0
	window.SetFullscreen(SDL_WINDOW_WINDOWED)
	window.SetPosition(sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED)
	canvasWidth, canvasHeight = windowWidth, windowHeight
	Clear()
	isFullscreen = false
	return nil
}

func SetVirtualSize(w, h int32) error {
	vWidth, vHeight = w, h
	screenRenderer.SetLogicalSize(vWidth, vHeight)
	err := Clear();	if err != nil {return errors.Wrap(err, "could not clear window after changing virtual size")}
	return nil
}
func SetWindowSize(w, h int32) error {
	windowWidth, windowHeight = w, h
	window.SetSize(w, h)
	return nil
}

func FillScreen(r, g, b, a uint8) error {
	err := screenRenderer.SetDrawColor(r, g, b, a);	if err != nil {return errors.Wrap(err, "could not set draw color for fill operation")}
	err = screenRenderer.FillRect(nil);	if err != nil {return errors.Wrap(err, "could not execute fill operation")}
	return nil
}
func Clear() error {
	FillScreen(0, 0, 0, 0)
	return nil
}
func Present() error {
	screenRenderer.Present()
	frameCount += 1
	return nil
}
