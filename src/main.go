package main

//import "github.com/veandco/go-sdl2/sdl"
import "./renderEngine"
import "fmt"

func main() {
	r := renderEngine.CreateRenderEngine(true)
    defer r.Destroy()
    fmt.Println(r)
}
