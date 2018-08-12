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
var Buttons mouseButtons

func init() {
	Buttons = mouseButtons{
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
			Buttons.Left:   RELEASED.copy(),
			Buttons.Center: RELEASED.copy(),
			Buttons.Right:  RELEASED.copy(),
			Buttons.X1:     RELEASED.copy(),
			Buttons.X2:     RELEASED.copy(),
		},
	}
}
