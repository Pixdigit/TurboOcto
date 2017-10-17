package main

import "github.com/veandco/go-sdl2/sdl"
import "./GameEngine"
import "fmt"

func main() {

	conf := GameEngine.CreateDefaultConfiguration()
	env := GameEngine.CreateEnvironment(conf)
    defer env.Destroy()
	sprite := GameEngine.LoadSphericalObject("./assets/images/test1.bmp", env)

	printer := func(key sdl.Scancode) {fmt.Println(sdl.GetKeyName(sdl.GetKeyFromScancode(key)))}
    GameEngine.RegisterKeyPressHandler(printer)
	running := true

    for running {
		GameEngine.Blit(sprite, env)
		GameEngine.Flip(env)
        GameEngine.Update()
    }
    sdl.Delay(1000)
}
