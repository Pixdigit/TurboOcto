package TurboOcto

import "github.com/veandco/go-sdl2/sdl"

var renderer *sdl.Renderer
var window *sdl.Window

func createSurface() {
    windowFlags := uint32(sdl.WINDOW_SHOWN) | uint32(sdl.WINDOW_FULLSCREEN_DESKTOP)
    window, _ = sdl.CreateWindow("", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 0, 0, windowFlags)
    renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_PRESENTVSYNC)
    Clear()
}

//TODO: Expand to include FavIcon
//Yet waiting for implementation of Sprite system
func SetTitle(title string) {
    window.SetTitle(title)
}

func Fullscreen() {
    //TODO: Wait until go-sdl2 is fixed to not require displayMode
    displayIndex, _ := window.GetDisplayIndex()
    var displayMode sdl.DisplayMode
    dmode := & displayMode
    sdl.GetDesktopDisplayMode(displayIndex, dmode)
    w, h := dmode.W, dmode.H
    window.SetSize(w, h)
    window.SetFullscreen(sdl.WINDOW_FULLSCREEN)
    Fill(0, 0, 0, 0)
    Clear()
}
func Windowed(w, h int32) {
    const SDL_WINDOW_WINDOWED = 0
    window.SetFullscreen(SDL_WINDOW_WINDOWED)
    window.SetSize(w, h)
    Clear()
}

func Fill(r, g, b, a uint8) {
    oldR, oldG, oldB, oldA, _ := renderer.GetDrawColor()
    renderer.SetDrawColor(r, g, b, a)
    renderer.FillRect(nil)
    renderer.SetDrawColor(oldR, oldG, oldB, oldA)
}
func Clear() {
    Fill(0, 0, 0, 0)
    Present()
}
func Present() {
    renderer.Present()
}
