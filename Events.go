package turboOcto

import (
	"github.com/veandco/go-sdl2/sdl"
	"gitlab.com/Pixdigit/geometry"
)

type MouseState struct {
	Pos             geometry.Point
	Movement        geometry.Vector
	Scroll          geometry.Scalar
	ScrollRelative  geometry.Scalar
	ButtonsHeld     []bool
	ButtonsClicked  []bool
	ButtonsReleased []bool
}

var Mouse MouseState

const (
	LEFT = iota
	CENTER
	RIGHT
	X1
	X2
)

func init() {
	Mouse = MouseState{
		geometry.Point{0, 0},
		geometry.Vector{0, 0},
		0,
		0,
		make([]bool, 3),
		make([]bool, 3),
		make([]bool, 3),
	}
}

func UpdateEvents() {
	//Reset frame dependend variables
	Mouse.Movement.X = 0
	Mouse.Movement.Y = 0
	Mouse.ScrollRelative = 0
	for i := range Mouse.ButtonsClicked {
		Mouse.ButtonsClicked[i] = false
	}
	for i := range Mouse.ButtonsReleased {
		Mouse.ButtonsReleased[i] = false
	}

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.MouseButtonEvent:
			pos := geometry.Point{
				geometry.Scalar(e.X),
				geometry.Scalar(e.Y),
			}
			Mouse.Pos.MoveTo(pos)
			if e.Type == sdl.MOUSEBUTTONDOWN {
				switch e.Button {
				case sdl.BUTTON_LEFT:
					setMouseButtonState(LEFT, true)
				case sdl.BUTTON_MIDDLE:
					setMouseButtonState(CENTER, true)
				case sdl.BUTTON_RIGHT:
					setMouseButtonState(RIGHT, true)
				case sdl.BUTTON_X1:
					setMouseButtonState(X1, true)
				case sdl.BUTTON_X2:
					setMouseButtonState(X2, true)
				}
			} else if e.Type == sdl.MOUSEBUTTONUP {
				switch e.Button {
				case sdl.BUTTON_LEFT:
					setMouseButtonState(LEFT, false)
				case sdl.BUTTON_MIDDLE:
					setMouseButtonState(CENTER, false)
				case sdl.BUTTON_RIGHT:
					setMouseButtonState(RIGHT, false)
				case sdl.BUTTON_X1:
					setMouseButtonState(X1, false)
				case sdl.BUTTON_X2:
					setMouseButtonState(X2, false)
				}
			}
		case *sdl.MouseMotionEvent:
			//TODO: Change relative motion too
			pos := geometry.Point{
				geometry.Scalar(e.X),
				geometry.Scalar(e.Y),
			}
			Mouse.Pos.MoveTo(pos)
			Mouse.Movement = geometry.Vector{
				geometry.Scalar(e.XRel),
				geometry.Scalar(e.YRel),
			}
		case *sdl.MouseWheelEvent:
			Mouse.ScrollRelative = geometry.Scalar(e.X)
			Mouse.Scroll += geometry.Scalar(e.X)
		}
	}
}
func setMouseButtonState(buttonIndex int32, isDown bool) {
	for len(Mouse.ButtonsClicked)-1 < int(buttonIndex) {
		Mouse.ButtonsHeld = append(Mouse.ButtonsHeld, false)
		Mouse.ButtonsClicked = append(Mouse.ButtonsClicked, false)
		Mouse.ButtonsReleased = append(Mouse.ButtonsReleased, false)
	}

	if Mouse.ButtonsHeld[buttonIndex] {
		if !isDown {
			Mouse.ButtonsReleased[buttonIndex] = true
			Mouse.ButtonsHeld[buttonIndex] = false
		}
	} else {
		if isDown {
			Mouse.ButtonsClicked[buttonIndex] = true
			Mouse.ButtonsHeld[buttonIndex] = true
		}
	}
}
