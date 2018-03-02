package main

import (
    "github.com/Pixdigit/TurboOcto"
    "github.com/veandco/go-sdl2/sdl"
)

func main() {
    TurboOcto.SetTitle("TEST")
    sdl.Delay(3000)
    TurboOcto.Windowed(500, 500)
    sdl.Delay(3000)
    TurboOcto.Fullscreen()
    sdl.Delay(300)
    for i:=0; i < 5; i++ {
        TurboOcto.Fill(255, 0, 0, 255)
        TurboOcto.Present()
        sdl.Delay(300)
        TurboOcto.Fill(0, 255, 0, 255)
        TurboOcto.Present()
        sdl.Delay(300)
        TurboOcto.Fill(0, 0, 255, 255)
        TurboOcto.Present()
        sdl.Delay(300)
    }
TurboOcto.Quit()
}
