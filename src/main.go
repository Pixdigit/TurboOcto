package main

import "github.com/veandco/go-sdl2/sdl"
import "./renderEngine"
import "fmt"

func main() {
	r := renderEngine.CreateRenderEngine(true)
    image := r.LoadImage("./assets/images/test1.bmp")
    r.Blit(image, &sdl.Rect{0, 0, 200, 200})
    r.Flip()
    defer r.Destroy()
    fmt.Println(r)
    sdl.Delay(5000)
}
