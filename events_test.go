package turboOcto

import (
	"testing"
	_ "time"

	"gitlab.com/Pixdigit/geometry"
	tools "gitlab.com/Pixdigit/goTestTools"
)

func TestEventScaling(t *testing.T) {
	t.Skip("test skipped due to unsable behaviour")
	Windowed()
	SetWindowSize(geometry.Size{20, 20})
	SetVirtualSize(geometry.Size{50, 50})
	window.WarpMouseInWindow(10, 10)
	updateEvents()
	t.Logf("Mouse at: %+v\n", Mouse.Pos)
	t.Log("This error is often caused by some timing delay. So in most cases you can ignore it.")
	tools.Test(Mouse.Pos.X == 10 && Mouse.Pos.Y == 10, "event handler did not scale input", t)
}

func TestInteractive(t *testing.T) {
	//TODO: Wait for frame manipulation sub library
	t.Log("This test is skipped for now")
	t.SkipNow()

	/*
		if testing.Short() {
			t.SkipNow()
		}

		SetWindowSize(geometry.Size{500, 500})
		SetVirtualSize(geometry.Size{500, 500})
		Windowed()
		window.SetTitle("Click all white boxes")
		Clear()
		Render()

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
			Render()
			Clear()
			//sdl.Delay(14)
		}
		time.Sleep(1 * time.Second)
	*/
}
