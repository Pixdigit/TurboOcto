package turboOcto

import "github.com/veandco/go-sdl2/sdl"

type keyboardState map[sdl.Scancode]*buttonState

var Keyboard keyboardState

func init() {
	Keyboard = make(map[sdl.Scancode]*buttonState)
}
