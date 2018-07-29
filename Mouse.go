package turboOcto

import (
	"github.com/veandco/go-sdl2/sdl"
	"gitlab.com/Pixdigit/geometry"
)

type buttonPosition uint8

type mouseState struct {
	Pos            geometry.Point
	Movement       geometry.Vector
	Scroll         geometry.Scalar
	ScrollRelative geometry.Scalar
	Buttons        map[buttonPosition]*buttonState
}

type mouseButtons struct {
	Left   buttonPosition
	Center buttonPosition
	Right  buttonPosition
	X1     buttonPosition
	X2     buttonPosition
}

var Mouse mouseState
var ButtonTypes mouseButtons

func init() {
	ButtonTypes = mouseButtons{
		Left:   buttonPosition(sdl.BUTTON_LEFT),
		Center: buttonPosition(sdl.BUTTON_MIDDLE),
		Right:  buttonPosition(sdl.BUTTON_RIGHT),
		X1:     buttonPosition(sdl.BUTTON_X1),
		X2:     buttonPosition(sdl.BUTTON_X2),
	}
	Mouse = mouseState{
		geometry.Point{0, 0},
		geometry.Vector{0, 0},
		0,
		0,
		map[buttonPosition]*buttonState{
			ButtonTypes.Left:   &RELEASED,
			ButtonTypes.Center: &RELEASED,
			ButtonTypes.Right:  &RELEASED,
			ButtonTypes.X1:     &RELEASED,
			ButtonTypes.X2:     &RELEASED,
		},
	}
}
