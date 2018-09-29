package turboOcto

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
	geo "gitlab.com/Pixdigit/geometry"
	"gitlab.com/Pixdigit/simpleTimer"
	"gitlab.com/Pixdigit/uniqueID"
)

//Sprites should only be initialized with NewSprite or the LoadSprite[…] functions
type Sprite struct {
	*Rect
	AllowFrameSkipping bool
	FrameIndex         int
	animationStatus    Runlevel
	delays             []int
	frames             []*Frame
	id                 uniqueID.ID
	timer              simpleTimer.Timer
	TimerMode          timerMode
}

type timerMode int

const (
	USE_FRAME_COUNT = timerMode(iota)
	USE_TIME_PASSED
)

var sprites []*Sprite

func NewSprite() (*Sprite, error) {
	newSprite := &Sprite{}

	newSprite.id = uniqueID.NewID()

	newSprite.TimerMode = timerMode(Cfg.DefaultSpriteTimerMode)
	newSprite.AllowFrameSkipping = Cfg.AllowFrameSkipping
	newSprite.animationStatus = STOPPED
	//TODO: sort out positioning
	//Create Rect. dimensions dont matter
	rect, err := NewRectFromGeometryRect(geo.NewRect(geo.Point{0, 0}, geo.Size{0, 0}, geo.CENTER));	if err != nil {return nil, errors.Wrap(err, "could not create geometry for sprite")}
	newSprite.Rect = rect

	frame, err := NewEmptyFrame();	if err != nil {return nil, errors.Wrap(err, "could not create empty Frame for new Sprite")}
	newSprite.frames = []*Frame{frame}

	return newSprite, nil
}

func (s *Sprite) ID() uniqueID.ID {
	return s.id
}

func (s *Sprite) render() error {
	//TODO: check when implementing visibility
	// if !s.Visible {
	// 	return nil
	// }

	err := s.update();	if err != nil {return errors.Wrap(err, "could not update sprite")}
	//TODO: Figure out positioning and return error
	_ = s.frames[s.FrameIndex].render()

	return nil
}

func (s *Sprite) SetDelay(time int) error {

	s.delays[s.FrameIndex] = time
	ok := s.validateDelays()
	if !ok {
		return errors.New("Sprite does not have any waiting time and will be blitted inifinitly")
	}

	return nil
}

func (s *Sprite) validateDelays() bool {
	cummulativeWaitingTime := 0
	for _, delay := range s.delays {
		cummulativeWaitingTime += delay
	}
	if cummulativeWaitingTime == 0 && s.AllowFrameSkipping {
		s.Stop()
		return false
	} else {
		return true
	}
}

func (s *Sprite) update() error {
	//Update the timer
	if s.animationStatus == RUNNING {
		if s.TimerMode == USE_FRAME_COUNT {
			s.timer.Update(float64(frameCount))
		} else if s.TimerMode == USE_TIME_PASSED {
			currentTime := sdl.GetTicks()
			s.timer.Update(float64(currentTime))
		}
	} else if s.animationStatus == STOPPED {
		return nil
	}

	//Update the frame index
	incrementFrameIndex := func() {
		s.timer.CarryReset()
		s.FrameIndex = (s.FrameIndex + 1) % len(s.frames)
		s.timer.Duration = float64(s.delays[s.FrameIndex])
		s.setTimerStartOffset()
	}
	if s.timer.Ended() {
		if s.AllowFrameSkipping {
			//Update frame index until timer is below delay
			for s.timer.Ended() {
				incrementFrameIndex()
			}
		} else {
			incrementFrameIndex()
		}
	}

	// TODO: Implement Visible. Note that framecount should not be incremented
	// TODO: Update frame position
	// s.Rect.SetSize(s.dimensions[s.FrameIndex])

	return nil
}

func (s *Sprite) setTimerStartOffset() {
	if s.TimerMode == USE_FRAME_COUNT {
		s.timer.LastUpdate = float64(frameCount)
	} else if s.TimerMode == USE_TIME_PASSED {
		currentTime := sdl.GetTicks()
		s.timer.LastUpdate = float64(currentTime)
	}
}

func (s *Sprite) Start() error {
	s.animationStatus = RUNNING
	s.setTimerStartOffset()
	return nil
}
func (s *Sprite) Stop() error {
	s.timer.Reset()
	s.FrameIndex = 0
	s.animationStatus = STOPPED
	return nil
}
func (s *Sprite) Pause() error {
	s.animationStatus = PAUSED
	return nil
}
