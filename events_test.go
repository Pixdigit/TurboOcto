package turboOcto

import (
	"github.com/veandco/go-sdl2/sdl"
	"testing"
)

func TestEventScaling(t *testing.T) {
	e := &sdl.MouseButtonEvent{
		Type:      sdl.MOUSEBUTTONDOWN,
		Timestamp: 1337,
		WindowID:  0,
		Button:    sdl.BUTTON_MIDDLE,
		X:         1337,
		Y:         1337,
		State:     sdl.PRESSED}
	filtered, err := sdl.PushEvent(e)
	if err != nil {
		wrapErr(err, "could not push test event", t)
	}
	test(!filtered, "test event was not pushed succesfully to the queue", t)
	Windowed(20, 20)
	UpdateEvents()
	test(Mouse.X < 20 && Mouse.X >= 0, "event handler did not scale input", t)
	test(Mouse.X != 0, "mouse position did not change on mouse event", t)

}
