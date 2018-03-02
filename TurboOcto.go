package TurboOcto

import "github.com/veandco/go-sdl2/sdl"

func init() {
    sdl.Init(sdl.INIT_EVERYTHING)
    createSurface()
}

func Quit() {
    renderer.Destroy()
    window.Destroy()
    sdl.Quit()
}
