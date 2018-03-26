package TurboOcto

import "github.com/veandco/go-sdl2/sdl"

var Events []sdl.Event

func UpdateEvents() []sdl.Event {
    Events = []sdl.Event{}
    for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
/*        switch e := event.(type) {
        case sdl.MouseButtonEvent:
            x, y := scalePoint(e.X, e.Y)
            e.X, e.Y = x, y
        }*/
        Events = append(Events, event)
    }
    return Events
}
func scalePoint(x, y int32) (int32, int32) {
    if sizer == UNDERFIT_SCALE {
        xScale := float64(drawWidth) / float64(screenWidth)
        x = int32(float64(xOffset) + xScale * float64(x))
        yScale := float64(drawHeight) / float64(screenHeight)
        y = int32(float64(yOffset) + yScale * float64(y))
    }
    return x, y
}
