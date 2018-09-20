package turboOcto

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
	"gitlab.com/Pixdigit/geometry"
	tools "gitlab.com/Pixdigit/goTestTools"
)

func TestScreen(t *testing.T) {
	Windowed()
	testDelay := uint32(30)
	tools.TestAgainstStrings(
		func(s string) error { return SetDecoration(s, "assets/images/testIcon.png") },
		func() (string, error) { s := window.GetTitle(); return s, nil },
		"window title not set properly", t)
	sdl.Delay(testDelay)
	SetWindowSize(geometry.Size{500, 500})
	x, y := window.GetSize()
	tools.Test((x == 500 && y == 500), "window has not changed resolution correctly", t)
	sdl.Delay(testDelay)
	Fullscreen()
	//TODO: Implement test. SDL has no GetFullscreen
}

func TestRenderer(t *testing.T) {
	/*pixelFormat, _ := window.GetPixelFormat() //TODO: CHeck out error

	testPixels := func (r, g, b, a uint8) {
	    var pixels unsafe.Pointer
	    pitch := 100 //for this test pitch can be arbitrary
	    screenRenderer.ReadPixels(nil, pixelFormat, pixels, pitch)

	    pixelsa := (*[]uint32)(pixels)
	    for _, pixel := range(*pixelsa) {
	        print(pixel)
	    }
	}*/ //TODO: Wait for go-sdl2 to return premade pixels array

	Windowed()
	SetWindowSize(geometry.Size{500, 500})
	FillScreen(255, 127, 0, 255)
	sdl.Delay(50)
	baseFrameCount := frameCount
	Render()
	tools.Test(frameCount-baseFrameCount == 1, "rendering did not increase frame count", t)
	if !testing.Short() {
		sdl.Delay(500)
		FillScreen(0, 255, 127, 255)
		sdl.Delay(50)
		Render()
		tools.Test(frameCount-baseFrameCount == 2, "rendering did not increase frame count", t)
		sdl.Delay(500)
		FillScreen(127, 0, 255, 255)
		sdl.Delay(50)
		Render()
		tools.Test(frameCount-baseFrameCount == 3, "rendering did not increase frame count", t)
		sdl.Delay(500)
	}
}
