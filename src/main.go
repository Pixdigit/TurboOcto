package main

import "github.com/veandco/go-sdl2/sdl"
import "github.com/ByteArena/box2d"
import "./GameEngine"
import "fmt"

func main() {

	conf := GameEngine.CreateDefaultCOnfiguration()
	env := GameEngine.CreateEnvironment(conf)
    defer env.Destroy()
	fmt.Println(env == box2d.B2World{})
	/*sprite := GameEngine.LoadSphericalObject("./assets/images/test1.bmp", env)
    GameEngine.Blit(sprite, env)
    GameEngine.Flip(env)
    printer := func(key sdl.Scancode) {fmt.Println(sdl.GetKeyName(sdl.GetKeyFromScancode(key)))}
    GameEngine.RegisterKeyPressHandler(printer)
	running := true
    for running {
        GameEngine.Update()
    }*/
    sdl.Delay(1000)
}
