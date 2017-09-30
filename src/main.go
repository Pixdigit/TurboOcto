package main

import "github.com/veandco/go-sdl2/sdl"
import "./GameEngine"
import "fmt"

func main() {
	r := GameEngine.CreateGraphicsInterface(true)
    defer r.Destroy()
    image := r.LoadImage("./assets/images/test1.bmp")
    r.Blit(image, &sdl.Rect{0, 0, 200, 200})
    r.Flip()
    printer := func(key sdl.Scancode) {fmt.Println(sdl.GetKeyName(sdl.GetKeyFromScancode(key)))}
    GameEngine.RegisterKeyPressHandler(printer)
    running := true
    for running {
        running = GameEngine.Update()
    }
    sdl.Delay(1000)
}
