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
	delays             []int
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
	//TODO: check when implementing visibility
	// if !s.Visible {
	// 	return nil
	// }

	err := s.update();	if err != nil {return errors.Wrap(err, "could not update sprite")}
	//TODO: Figure out positioning and return error
	err = s.Renderables[s.FrameIndex].render();	if err != nil {return errors.Wrap(err, "error during rendering sprite frame")}

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
	incrementFrameIndex := func() error {
		s.timer.CarryReset()
        //increment and wrao around
		s.FrameIndex = (s.FrameIndex + 1) % len(s.Renderables)
		for len(s.Renderables) != len(s.delays) {
			if len(s.Renderables) > len(s.delays) {
				s.delays = append(s.delays, 0)
                ok := s.validateDelays()
    			if !ok {
                    //revert change and return error
                    s.delays = s.delays[:len(s.delays)-1]
    				return errors.New("changing frame count lead to invalid delays")
    			}
			} else {
                lastDelay := s.delays[len(s.delays)-1]
				s.delays = s.delays[:len(s.delays)-1]
                ok := s.validateDelays()
                if !ok {
                    //revert change and return error
                    s.delays = append(s.delays, lastDelay)
    				return errors.New("changing frame count lead to invalid delays")
    			}

			}
		}
		s.timer.Duration = float64(s.delays[s.FrameIndex])
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

	// TODO: Implement Visible. Note that framecount should not be incremented

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
