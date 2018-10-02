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
	*Collection
	AllowFrameSkipping bool
	FrameIndex         int
	animationStatus    Runlevel
	Delays             []int
	id                 uniqueID.ID
	timer              simpleTimer.Timer
	TimerMode          timerMode
	Visible            bool
}

type timerMode int

const (
	USE_FRAME_COUNT = timerMode(iota)
	USE_TIME_PASSED
)

var sprites []*Sprite

func NewSprite() (*Sprite, error) {
	//The Renderables have to be initialized
	newSprite := &Sprite{
		Collection: &Collection{
			Renderables: make([]PositionedRenderable, 1),
			anchor:      &geo.Point{0, 0},
		},
	}

	newSprite.id = uniqueID.NewID()

	newSprite.TimerMode = timerMode(Cfg.DefaultSpriteTimerMode)
	newSprite.AllowFrameSkipping = Cfg.AllowFrameSkipping
	newSprite.animationStatus = STOPPED

	frame, err := NewEmptyFrame();	if err != nil {return nil, errors.Wrap(err, "could not create empty Frame for new Sprite")}
	newSprite.Renderables = []PositionedRenderable{frame}

	return newSprite, nil
}

func (s *Sprite) ID() uniqueID.ID {
	return s.id
}

func (s *Sprite) render() error {

	err := s.update();	if err != nil {return errors.Wrap(err, "could not update sprite")}
	if s.Visible {
		err = s.Renderables[s.FrameIndex].render();	if err != nil {return errors.Wrap(err, "error during rendering sprite frame")}
	}
	return nil
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
	incrementFrameIndex := func() error {
		s.timer.CarryReset()
		//increment and wrap around
		s.FrameIndex = (s.FrameIndex + 1) % len(s.Renderables)
		for len(s.Renderables) != len(s.Delays) {
			if len(s.Renderables) > len(s.Delays) {
				s.Delays = append(s.Delays, 0)
				ok := s.validateDelays()
				if !ok {
					//revert change and return error
					s.Delays = s.Delays[:len(s.Delays)-1]
					return errors.New("changing frame count lead to invalid delays")
				}
			} else {
				lastDelay := s.Delays[len(s.Delays)-1]
				s.Delays = s.Delays[:len(s.Delays)-1]
				ok := s.validateDelays()
				if !ok {
					//revert change and return error
					s.Delays = append(s.Delays, lastDelay)
					return errors.New("changing frame count lead to invalid delays")
				}

			}
		}
		s.timer.Duration = float64(s.Delays[s.FrameIndex])
		s.setTimerStartOffset()
		return nil
	}
	if s.timer.Ended() {
		if s.AllowFrameSkipping {
			//Update frame index until timer is below delay
			for s.timer.Ended() {
				err := incrementFrameIndex();	if err != nil {return errors.Wrap(err, "error while incrementing frame index with skipping")}
			}
		} else {
			err := incrementFrameIndex();	if err != nil {return errors.Wrap(err, "error while incrementing frame index without skipping")}
		}
	}

	return nil
}

func (s *Sprite) SetDelay(time int) error {

	s.Delays[s.FrameIndex] = time
	ok := s.validateDelays()
	if !ok {
		return errors.New("Sprite does not have any waiting time and will be blitted inifinitly")
	}

	return nil
}

func (s *Sprite) validateDelays() bool {
	cummulativeWaitingTime := 0
	for _, delay := range s.Delays {
		cummulativeWaitingTime += delay
	}
	if cummulativeWaitingTime == 0 && s.AllowFrameSkipping {
		s.Stop()
		return false
	} else {
		return true
	}
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
