package TurboOcto

import (
    "github.com/veandco/go-sdl2/sdl"
    "testing"
    "os"
)

func TestMain(m *testing.M) {
    createScreen()
    result := m.Run()
    Quit()
    os.Exit(result)
}

func TestScreen(t *testing.T) {
    testDelay := uint32(30)
    testAgainstStrings(SetTitle, window.GetTitle, "window title remains unchanged", t)
    sdl.Delay(testDelay)
    Windowed(500, 500)
    x, y := window.GetSize()
    test((x == 500 && y == 500), "window has not changed resolution correctly", t)
    sdl.Delay(testDelay)
    Fullscreen()
    //TODO: Implement test. SDL has no GetFullscreen
}

func TestRenderer(t *testing.T) {
    /*pixelFormat, _ := window.GetPixelFormat() //TODO: CHeck out error

    testPixels := func (r, g, b, a uint8) {
        var pixels unsafe.Pointer
        pitch := 100 //for this test pitch can be arbitrary
        renderer.ReadPixels(nil, pixelFormat, pixels, pitch)

        pixelsa := (*[]uint32)(pixels)
        for _, pixel := range(*pixelsa) {
            print(pixel)
        }
    }*/ //TODO: Wait for go-sdl2 to return premade pixels array

    FillScreen(255, 127, 0, 255)
    Present()
    FillScreen(0, 255, 127, 255)
    Present()
    FillScreen(127, 0, 255, 255)
    Present()
}