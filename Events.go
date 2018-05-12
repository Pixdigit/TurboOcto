package TurboOcto

import (
	"github.com/veandco/go-sdl2/sdl"
)

type MouseState struct {
	X, Y            int32
	XRel, YRel      int32
	Scrolled        int32
	ScrollRelative  int32
	ButtonsHeld     []bool
	ButtonsClicked  []bool
	ButtonsReleased []bool
}

var Mouse MouseState

func init() {
	Mouse = MouseState{
		0, 0, 0, 0, 0, 0,
		make([]bool, 3),
		make([]bool, 3),
		make([]bool, 3),
	}
}

func UpdateEvents() {
	//Reset frame dependend variables
	Mouse.XRel, Mouse.YRel = 0, 0
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
			x, y := scalePoint(e.X, e.Y)
			Mouse.X, Mouse.Y = x, y
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
			x, y := scalePoint(e.X, e.Y)
			Mouse.X, Mouse.Y = x, y
			Mouse.XRel, Mouse.YRel = e.XRel, e.YRel
		case *sdl.MouseWheelEvent:
			x, y := scalePoint(e.X, e.Y)
			Mouse.X, Mouse.Y = x, y
			Mouse.ScrollRelative = e.X
			Mouse.Scrolled += e.X
		}
	}
}
func scalePoint(x, y int32) (int32, int32) {
	if sizer == UNDERFIT_SCALE {
		xScale := float64(drawWidth) / float64(screenWidth)
		x = int32(float64(xOffset) + xScale*float64(x))
		yScale := float64(drawHeight) / float64(screenHeight)
		y = int32(float64(yOffset) + yScale*float64(y))
	}
	return x, y
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
