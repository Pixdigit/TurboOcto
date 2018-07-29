package turboOcto

import (
	"github.com/veandco/go-sdl2/sdl"
	"gitlab.com/Pixdigit/geometry"
)

func UpdateEvents() {
	//Reset frame dependend variables
	Mouse.Movement.X = 0
	Mouse.Movement.Y = 0
	Mouse.ScrollRelative = 0
	for i := range Mouse.Buttons {
		Mouse.Buttons[i].Changed = false
	}
	for i := range Keyboard {
		Keyboard[i].Changed = false
	}

	//Process events
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.MouseButtonEvent:
			newPos := geometry.Point{
				geometry.Scalar(e.X),
				geometry.Scalar(e.Y),
			}
			Mouse.Pos.MoveTo(newPos)
			if e.Type == sdl.MOUSEBUTTONDOWN {
				Mouse.Buttons[buttonPosition(e.Button)].update(true)
			} else if e.Type == sdl.MOUSEBUTTONUP {
				Mouse.Buttons[buttonPosition(e.Button)].update(false)
			}
		case *sdl.MouseMotionEvent:
			newPos := geometry.Point{
				geometry.Scalar(e.X),
				geometry.Scalar(e.Y),
			}
			Mouse.Pos.MoveTo(newPos)
			Mouse.Movement = geometry.Vector{
				geometry.Scalar(e.XRel),
				geometry.Scalar(e.YRel),
			}
		case *sdl.MouseWheelEvent:
			Mouse.ScrollRelative = geometry.Scalar(e.X)
			Mouse.Scroll += geometry.Scalar(e.X)
		case *sdl.KeyboardEvent:
			//Create entry for key if it does not exist yet
			_, ok := Keyboard[e.Keysym.Scancode]
			if !ok {
				Keyboard[e.Keysym.Scancode] = &RELEASED
			}
			//Update key
			if e.Type == sdl.KEYDOWN {
				Keyboard[e.Keysym.Scancode].update(true)
			} else if e.Type == sdl.KEYUP {
				Keyboard[e.Keysym.Scancode].update(false)
			}
		}
	}
}
