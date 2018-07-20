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
			Mouse.Pos.MoveTo(scalePoint(pos))
			if e.Type == sdl.MOUSEBUTTONDOWN {
				switch e.Button {
				case sdl.BUTTON_LEFT:
					setMouseButtonState(1, true)
				case sdl.BUTTON_MIDDLE:
					setMouseButtonState(2, true)
				case sdl.BUTTON_RIGHT:
					setMouseButtonState(3, true)
				case sdl.BUTTON_X1:
					setMouseButtonState(4, true)
				case sdl.BUTTON_X2:
					setMouseButtonState(5, true)
				}
			} else if e.Type == sdl.MOUSEBUTTONUP {
				switch e.Button {
				case sdl.BUTTON_LEFT:
					setMouseButtonState(1, false)
				case sdl.BUTTON_MIDDLE:
					setMouseButtonState(2, false)
				case sdl.BUTTON_RIGHT:
					setMouseButtonState(3, false)
				case sdl.BUTTON_X1:
					setMouseButtonState(4, false)
				case sdl.BUTTON_X2:
					setMouseButtonState(5, false)
				}
			}
		case *sdl.MouseMotionEvent:
			//TODO: Change relative motion too
			pos := geometry.Point{
				geometry.Scalar(e.X),
				geometry.Scalar(e.Y),
			}
			Mouse.Pos.MoveTo(scalePoint(pos))
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
func scalePoint(p geometry.Point) geometry.Point {
	if sizer == UNDERFIT_SCALE {
		xScale := float64(drawWidth) / float64(screenWidth)
		p.X = geometry.Scalar(float64(xOffset) + xScale*float64(p.X))
		yScale := float64(drawHeight) / float64(screenHeight)
		p.Y = geometry.Scalar(float64(yOffset) + yScale*float64(p.Y))
	}
	return p
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
