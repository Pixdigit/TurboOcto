package turboOcto

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
	"gitlab.com/Pixdigit/geometry"
)

type sprite struct {
	frames     []*sdl.Texture
	dimensions []geometry.Size
	*geometry.Rect
	delays             []int32
	animationStatus    Runlevel
	timerMode          timerMode
	timer              int32
	lastBlit           int32
	lastFrameCount     int32
	lastTimer          int32
	AllowFrameSkipping bool
	FrameIndex         int32
	layer              int32
	Visible            bool
	constraint         func(*sprite) error
}

type timerMode int

const (
	USE_FRAME_COUNT = timerMode(iota)
	USE_TIME_PASSED
)

var sprites []*sprite

func NewSprite() (*sprite, error) {
	newSprite := &sprite{}
	//ensure sprite is the topmost of level 0
	sprites = append([]*sprite{newSprite}, sprites...)
	err := newSprite.ChangeLayer(0);	if err != nil {return &sprite{}, errors.Wrap(err, "could not change layer for new Sprite")}

	newSprite.timerMode = timerMode(Cfg.SpriteTimerMode)
	newSprite.AllowFrameSkipping = Cfg.AllowFrameSkipping
	newSprite.lastFrameCount = frameCount
	newSprite.animationStatus = STOPPED
	newSprite.Visible = true
	// TODO: Is this needed?
	//newSprite.dimensions = []Size{Size{}}
	rect := geometry.NewRect(geometry.Point{0, 0}, geometry.Size{0, 0}, geometry.CENTER)
	newSprite.Rect = &rect

	return newSprite, nil
}

//ChangeLayer is quite performance heavy since it moves an element within a slice
//this, however, will save computational power at Blit-Time
func (s *sprite) ChangeLayer(layer int32) error {
	s.layer = layer
	for i := len(sprites) - 1; i >= 0; i-- {
		sp := sprites[i]
		if sp.layer <= s.layer {
			for i, sp := range sprites {
				if s == sp {
					var newSprites []*sprite
					newSprites = append(newSprites, sprites[:i]...)
					newSprites = append(newSprites, sprites[i+1:]...)
					sprites = newSprites
				}
			}
			var newSprites []*sprite
			newSprites = append(newSprites, sprites[:i]...)
			newSprites = append(newSprites, s)
			newSprites = append(newSprites, sprites[i:]...)
			sprites = newSprites
			break
		}
	}
	return nil
}

func (s *sprite) SetDelay(time int32) error {
	cummulativeWaitingTime := int32(0)
	s.delays[s.FrameIndex] = time
	for _, delay := range s.delays {
		cummulativeWaitingTime += delay
	}
	if cummulativeWaitingTime == 0 && s.AllowFrameSkipping {
		s.Stop()
		return errors.New("Sprite does not have any waiting time and will be blitted inifinitly")
	}
	return nil
}

func (s *sprite) IncrementTime() error {
	var currentTime int32
	if s.animationStatus == RUNNING {
		currentTime := int32(sdl.GetTicks())
		if s.timerMode == USE_FRAME_COUNT {
			s.timer += frameCount - s.lastFrameCount
		} else if s.timerMode == USE_TIME_PASSED {
			s.timer += currentTime - s.lastBlit
		}
	} else if s.animationStatus == STOPPED {
		return nil
	}
	s.lastBlit = int32(currentTime)
	s.lastFrameCount = frameCount

	if s.timer >= s.delays[s.FrameIndex] {
		if s.AllowFrameSkipping {
			for s.timer >= s.delays[s.FrameIndex] {
				s.timer -= s.delays[s.FrameIndex]
				s.FrameIndex = (s.FrameIndex + 1) % int32(len(s.frames))
			}
		} else {
			//If we have no frame skipping ensure at least one blit
			s.timer = s.timer - s.delays[s.FrameIndex]
			if s.timer > s.lastTimer || (s.FrameIndex == 0 && s.timer == 0) {
				s.FrameIndex = (s.FrameIndex + 1) % int32(len(s.frames))
			}
		}
	}
	s.lastTimer = s.timer
	s.Rect.SetSize(s.dimensions[s.FrameIndex])
	return nil
}

func (s *sprite) SetConstraint(constraint func(*sprite) error) error {
	s.constraint = constraint
	s.constraint(s)
	return nil
}

func (s *sprite) SetSize(size geometry.Size) error {
	s.dimensions[s.FrameIndex] = size
	s.Rect.SetSize(size)
	return nil
}

func (s *sprite) BlitToScreen() error {
	if !s.Visible {
		return nil
	}
	size := s.Rect.Size()
	topLeft := s.Rect.PositionFrom(geometry.TOPLEFT)
	dstRect := &sdl.Rect{int32(topLeft.X), int32(topLeft.Y), int32(size.Width), int32(size.Height)}
	err := screenRenderer.Copy(s.frames[s.FrameIndex], nil, dstRect);	if err != nil {return errors.Wrap(err, "could not copy sprite frame to screenRenderer")}
	return nil
}

func (s *sprite) Blit(dstTexture *sdl.Texture) error {
	if !s.Visible {
		return nil
	}
	size := s.Rect.Size()
	topLeft := s.Rect.PositionFrom(geometry.TOPLEFT)
	dstRect := &sdl.Rect{int32(topLeft.X), int32(topLeft.Y), int32(size.Width), int32(size.Height)}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_TARGETTEXTURE);	if err != nil {return errors.Wrap(err, "could not create renderer to render to texture")}
	renderer.SetRenderTarget(dstTexture)
	err = renderer.Copy(s.frames[s.FrameIndex], nil, dstRect);	if err != nil {return errors.Wrap(err, "could not copy sprite frame to texture")}
	return nil
}

func (s *sprite) IsClicked(which buttonPosition) (bool, error) {
	return s.Rect.Contains(Mouse.Pos) && (*Mouse.Buttons[which] == PRESSING), nil
}

func (s *sprite) Start() error {
	s.animationStatus = RUNNING
	s.lastFrameCount = frameCount
	s.lastBlit = int32(sdl.GetTicks())
	s.lastTimer = s.timer - 1
	return nil
}
func (s *sprite) Stop() error {
	s.FrameIndex = 0
	s.timer = 0
	s.animationStatus = STOPPED
	return nil
}
func (s *sprite) Pause() error {
	s.animationStatus = PAUSED
	return nil
}

func updateAllSprites() error {
	for _, sp := range sprites {
		err := sp.IncrementTime();	if err != nil {return errors.Wrap(err, "could not increment time for all sprites")}
		err = sp.BlitToScreen();	if err != nil {return errors.Wrap(err, "could not blit all sprites")}
	}
	return nil
}
