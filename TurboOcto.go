package TurboOcto

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/pkg/errors"
    "fmt"
)

func init() {
    sdl.Init(sdl.INIT_EVERYTHING)
    err := initializeGraphics();    if err != nil {fmt.Println(errors.Wrap(err, "could not initialize graphics")); Quit()}
    err = initializeEnvironment();    if err != nil {fmt.Println(errors.Wrap(err, "could not initialize environment")); Quit()}
}

func Update() {
    Present()
    UpdateEvents()
}

func Quit() {
    renderer.Destroy()
    window.Destroy()
    sdl.Quit()
}
