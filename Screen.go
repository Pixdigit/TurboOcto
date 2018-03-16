package TurboOcto

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/pkg/errors"
)


type sizerType int32
type scalerType int32

var drawWidth, drawHeight int32
var currWidth, currHeight int32
var maxWidth, maxHeight int32
var xOffset, yOffset int32
var sizer sizerType
var scaler scalerType

const UNDERFIT_SCALE sizerType = 0
const OVERFIT_SCALE sizerType = 1
const STRECH_SCALE sizerType = 2
const FIX_SCALE sizerType = 3

//TODO implement scaling methods
const SIMPLE_SCALE scalerType = 1


func initializeGraphics() (err error) {
    windowFlags := uint32(sdl.WINDOW_SHOWN) | uint32(sdl.WINDOW_FULLSCREEN_DESKTOP)
    window, err = sdl.CreateWindow("", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 0, 0, windowFlags);    if err != nil {return errors.Wrap(err, "could not create window")}
    renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_PRESENTVSYNC);    if err != nil {return errors.Wrap(err, "could not create renderer")}
    drawWidth, drawHeight, err = renderer.GetOutputSize();    if err != nil {return errors.Wrap(err, "could not read output size")}
    currWidth, currHeight = drawWidth, drawHeight
    Clear()
    return nil
}

//TODO: Expand to include FavIcon
//    Yet waiting for implementation of Sprite system
func SetTitle(title string) error {
    window.SetTitle(title)
    return nil
}

func Fullscreen() {
    window.SetSize(maxWidth, maxHeight)
    window.SetFullscreen(sdl.WINDOW_FULLSCREEN)
    currWidth, currHeight = maxWidth, maxHeight
    FillScreen(0, 0, 0, 0)
    Clear()
}
func Windowed(w, h int32) {
    const SDL_WINDOW_WINDOWED = 0
    window.SetFullscreen(SDL_WINDOW_WINDOWED)
    window.SetSize(w, h)
    currWidth, currHeight = w, h
    window.SetPosition(sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED)
    Clear()
}

func SetSize(w, h int32) {
    drawWidth, drawHeight = w, h
}
func SetScaler(sizer sizerType, scaler scalerType) {
    sizer = sizer
    scaler = scaler
    Clear()

    switch sizer {
    case UNDERFIT_SCALE:
        aspectRatioWindow := float64(currWidth) / float64(currHeight)
        aspectRatioRenderer := float64(drawWidth) / float64(drawHeight)
        renderer.SetLogicalSize(drawWidth, drawHeight)
        if aspectRatioRenderer > aspectRatioWindow {
            //TODO: Implement test is offsets are correct
            //window is too thin horizontally
            xOffset = 0
            //                Get remaining height V          ScaleFactor V             Two sites V
            yOffset = int32(float64(currHeight) * (1 - float64(drawWidth) / float64(currWidth)) / 2)
        } else {
            //window is too small vertically
            renderer.SetLogicalSize(drawWidth, drawHeight)
            //               Get remaining height V           ScaleFactor V              Two sites V
            xOffset = int32(float64(currWidth) * (1 - float64(drawHeight) / float64(currHeight)) / 2)
            yOffset = 0
        }
    }

}


func FillScreen(r, g, b, a uint8) {
    oldR, oldG, oldB, oldA, _ := renderer.GetDrawColor()
    renderer.SetDrawColor(r, g, b, a)
    renderer.FillRect(nil)
    renderer.SetDrawColor(oldR, oldG, oldB, oldA)
}
func Clear() {
    FillScreen(0, 0, 0, 0)
}
func Present() {
    renderer.Present()
}
