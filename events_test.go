package turboOcto

// TODO:

/*
import (
	"testing"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"gitlab.com/Pixdigit/geometry"
	tools "gitlab.com/Pixdigit/goTestTools"
)

func TestEventScaling(t *testing.T) {
	Windowed()
	SetWindowSize(geometry.Size{20, 20})
	SetVirtualSize(geometry.Size{50, 50})
	e := &sdl.MouseButtonEvent{
		Type:      sdl.MOUSEBUTTONDOWN,
		Timestamp: 1337,
		WindowID:  0,
		Button:    sdl.BUTTON_MIDDLE,
		X:         10,
		Y:         10,
		State:     sdl.PRESSED}
	filtered, err := sdl.PushEvent(e);	if err != nil {tools.WrapErr(err, "could not push test event", t)}
	tools.Test(!filtered, "test event was not pushed succesfully to the queue", t)
	updateEvents()
	t.Logf("Mouse at: %+v\n", Mouse.Pos)
	tools.Test(Mouse.Pos.X == 10 && Mouse.Pos.Y == 10, "event handler did not scale input", t)
}

func TestInteractive(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	SetWindowSize(geometry.Size{500, 500})
	SetVirtualSize(geometry.Size{500, 500})
	Windowed()
	window.SetTitle("Click all white boxes")
	Clear()
	Present()

	testBox, _ := NewRect(geometry.Point{100, 100}, geometry.Size{100, 100}, geometry.TOPLEFT)
	testBox.Fill(255, 255, 255, 255)
	timeOut := time.After(5 * time.Second)
	run := true
	for run {
		select {
		case <-timeOut:
			run = false
			t.Error("Square was not clicked")
		default:
		}
		updateEvents()
		clicked, err := testBox.IsClicked(Buttons.Left);	if err != nil {tools.WrapErr(err, "could not check for clicked rect", t)}
		if clicked {
			run = false
			testBox.Fill(0, 255, 0, 255)
		} else {
			testBox.Fill(255, 255, 255, 255)
		}
		Present()
		Clear()
		//sdl.Delay(14)
	}
	time.Sleep(1 * time.Second)
}*/
