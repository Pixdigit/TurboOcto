package TurboOcto

import (
    "github.com/veandco/go-sdl2/sdl"
    "errors"
    "testing"
    "os"
)

var strings []string = []string{"TEST", "ẞönDérZäíſĉh€Ń", "1234567890", "", "\n"}

func TestMain(m *testing.M) {
    createScreen()
    result := m.Run()
    Quit()
    os.Exit(result)
}

func assert(success bool) error {
    if !success {
        return errors.New("assertion failed!")
    } else {
        return nil
    }
}

func test(success bool, errMsg string, t *testing.T) {
    err := assert(success)
    if err != nil {
        t.Error(errMsg)
    }
}

func testAgainstStrings(set func(s string)(error), get func()(string), errMsg string, t *testing.T) {
    for _, testString := range(strings) {
        set(testString)
        result := get()
        errorMsg := errMsg + ": failed at string \"" + testString + "\"; is " + result
        test(result == testString, errorMsg, t)
    }
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
