package TurboOcto

import "github.com/veandco/go-sdl2/sdl"

var Events []sdl.Event

func UpdateEvents() []sdl.Event {
    Events = []sdl.Event{}
    for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
        Events = append(Events, event)
    }
    return Events
}
