package eventEngine

import "github.com/veandco/go-sdl2/sdl"
import "fmt"


type EventEnv struct {
    newEvents []sdl.Event
}

func (env *EventEnv) UpdateEvents() (running bool, newEvents []sdl.Event) {
    fmt.Print("")
    running = true
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
        newEvents = append(newEvents, event)
        /*switch e := event.(type) {
		case *sdl.QuitEvent:
			running = false
            fmt.Println("QuitEvent")
		case *sdl.MouseMotionEvent:
			fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
				e.Timestamp, e.Type, e.Which, e.X, e.Y, e.XRel, e.YRel)
		case *sdl.MouseButtonEvent:
			fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
				e.Timestamp, e.Type, e.Which, e.X, e.Y, e.Button, e.State)
		case *sdl.MouseWheelEvent:
			fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
				e.Timestamp, e.Type, e.Which, e.X, e.Y)
        case *sdl.KeyDownEvent:
            append(newEvents, )
            if e.Keysym.Scancode == sdl.SCANCODE_ESCAPE {
                running = false
            }
		/*case *sdl.KeyUpEvent:
			fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
				e.Timestamp, e.Type, e.Keysym.Sym, e.Keysym.Mod, e.State, e.Repeat)
        case *sdl.ControllerButtonEvent:
            fmt.Printf("[%d ms] ControllerButton\ttype:%d\twhich:%c\tbutton:%d\tstate:%d\n",
                e.Timestamp, e.Type, e.Which, e.Button, e.State)
        case *sdl.ControllerDeviceEvent:
            var deviceEventType string
            switch e.Type {
            case sdl.CONTROLLERDEVICEADDED:
                deviceEventType = "Added"
            case sdl.CONTROLLERDEVICEREMOVED:
                deviceEventType = "Removed"
            case sdl.CONTROLLERDEVICEREMAPPED:
                deviceEventType = "Remapped"
            }
            fmt.Printf("[%d ms] ControllerDevice\ttype:%d\twhich:%c\n",
                e.Timestamp, deviceEventType, e.Which)
        */
        /*default:
			fmt.Printf("Some event\n")
        }*/
    }
    return running, newEvents
}