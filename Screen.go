package TurboOcto

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/pkg/errors"
)


type sizerType int32
type scalerType int32

var vWidth, vHeight int32
var drawWidth, drawHeight int32
var screenWidth, screenHeight int32
var xOffset, yOffset int32
var sizer sizerType
var scaler scalerType


//TODO: implement all scales
const UNDERFIT_SCALE sizerType = 0
const OVERFIT_SCALE sizerType = 1
const STRECH_SCALE sizerType = 2
const FIX_SCALE sizerType = 3

//TODO implement scaling methods
const SIMPLE_SCALE scalerType = 1

var renderer *sdl.Renderer
var window *sdl.Window
var displayIndex int //TODO: Dynamically update when window moved

var rmask uint32 = 0x000000ff;
var gmask uint32 = 0x0000ff00;
var bmask uint32 = 0x00ff0000;
var amask uint32 = 0xff000000;

func initializeGraphics() (err error) {
    //Default is LIL_ENDIAN
    if sdl.BYTEORDER == sdl.BIG_ENDIAN {
        rmask = 0xff000000;
        gmask = 0x00ff0000;
        bmask = 0x0000ff00;
        amask = 0x000000ff;
    }

    windowFlags := uint32(sdl.WINDOW_SHOWN) | uint32(sdl.WINDOW_FULLSCREEN_DESKTOP)
    window, err = sdl.CreateWindow("", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 0, 0, windowFlags);    if err != nil {return errors.Wrap(err, "could not create window")}

    renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_PRESENTVSYNC);    if err != nil {return errors.Wrap(err, "could not create renderer")}

    displayIndex, err := window.GetDisplayIndex();    if err != nil {return errors.Wrap(err, "could not get display index")}
    dmode, err := sdl.GetDesktopDisplayMode(displayIndex);    if err != nil {return errors.Wrap(err, "could not get display mode")}

    screenWidth, screenHeight = dmode.W, dmode.H
    drawWidth, drawHeight, err = renderer.GetOutputSize();    if err != nil {return errors.Wrap(err, "could not read output size")}
    vWidth, vHeight = drawWidth, drawHeight

    if ok, err := GetConf("fullscreen"); err != nil { return errors.Wrap(err, "could not get fullscreen conf")
    } else if ok.(bool) {
        Fullscreen()
    } else {
        Windowed(screenWidth / 4, screenHeight / 4)
    }

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
    window.SetSize(screenWidth, screenHeight)
    window.SetFullscreen(sdl.WINDOW_FULLSCREEN)
    drawWidth, drawHeight = screenWidth, screenHeight
    FillScreen(0, 0, 0, 0)
    Clear()
    SetConf("fullscreen", true)
}
func Windowed(w, h int32) {
    const SDL_WINDOW_WINDOWED = 0
    window.SetFullscreen(SDL_WINDOW_WINDOWED)
    window.SetSize(w, h)
    drawWidth, drawHeight = w, h
    window.SetPosition(sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED)
    Clear()
    SetConf("fullscreen", true)
}

func SetSize(w, h int32) {
    vWidth, vHeight = w, h
    renderer.SetLogicalSize(vWidth, vHeight)
}
func SetScaler(sizer sizerType, scaler scalerType) {
    sizer = sizer
    scaler = scaler
    Clear()

    switch sizer {
    case UNDERFIT_SCALE:
        aspectRatioWindow := float64(screenWidth) / float64(screenHeight)
        logicalAspectRatio := float64(vWidth) / float64(vHeight)
        renderer.SetLogicalSize(vWidth, vHeight)
        //More width than height
        if logicalAspectRatio > aspectRatioWindow {
            //TODO: Implement test is offsets are correct
            //window is too thin horizontally
            drawHeight = int32(float64(screenWidth) / float64(vWidth) * float64(screenHeight))
            drawWidth = screenWidth
            xOffset = 0
            yOffset = (screenHeight - drawHeight) / 2
        } else {
            //window is too small vertically
            drawHeight = screenHeight
            drawWidth = int32(float64(screenHeight) / float64(vHeight) * float64(screenWidth))
            xOffset = (screenWidth - drawWidth) / 2
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
