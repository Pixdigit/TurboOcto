package turboOcto

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
	tools "gitlab.com/Pixdigit/goTestTools"
)

func TestEventScaling(t *testing.T) {
	Windowed()
	SetWindowSize(20, 20)
	SetVirtualSize(50, 50)
	e := &sdl.MouseButtonEvent{
		Type:      sdl.MOUSEBUTTONDOWN,
		Timestamp: 1337,
		WindowID:  0,
		Button:    sdl.BUTTON_MIDDLE,
		X:         10,
		Y:         10,
		State:     sdl.PRESSED}
	filtered, err := sdl.PushEvent(e)
	if err != nil {
		tools.WrapErr(err, "could not push test event", t)
	}
	tools.Test(!filtered, "test event was not pushed succesfully to the queue", t)
	UpdateEvents()
	t.Logf("Mouse at: %+v\n", Mouse.Pos)
	tools.Test(Mouse.Pos.X == 10 && Mouse.Pos.Y == 10, "event handler did not scale input", t)

}
