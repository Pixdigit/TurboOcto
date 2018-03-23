package TurboOcto

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/img"
    "github.com/pkg/errors"
    "fmt"
)

func init() {
    sdl.Init(sdl.INIT_EVERYTHING)
    img.Init(0x0000000F)  // initialize all formats
    err := initializeEnvironment();    if err != nil {fmt.Println(errors.Wrap(err, "could not initialize environment")); Quit()}
    err = initializeGraphics();    if err != nil {fmt.Println(errors.Wrap(err, "could not initialize graphics")); Quit()}
}

func Update() {
    Present()
    UpdateEvents()
}

func Quit() {
    renderer.Destroy()
    window.Destroy()
    img.Quit()
    sdl.Quit()
}
