package eventEngine

import "github.com/veandco/go-sdl2/sdl"


type EventEnv struct {
    newEvents []sdl.Event
}

func (env *EventEnv) UpdateEvents() (newEvents []sdl.Event) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
        newEvents = append(newEvents, event)
    }
    return newEvents
}
